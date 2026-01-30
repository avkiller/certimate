package workflow

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/pocketbase/dbx"

	"github.com/certimate-go/certimate/internal/app"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/domain/dtos"
	"github.com/certimate-go/certimate/internal/workflow/dispatcher"
)

type WorkflowService struct {
	dispatcher dispatcher.WorkflowDispatcher

	workflowRepo    workflowRepository
	workflowRunRepo workflowRunRepository
	settingsRepo    settingsRepository
}

func NewWorkflowService(workflowRepo workflowRepository, workflowRunRepo workflowRunRepository, settingsRepo settingsRepository) *WorkflowService {
	srv := &WorkflowService{
		dispatcher: dispatcher.GetSingletonDispatcher(),

		workflowRepo:    workflowRepo,
		workflowRunRepo: workflowRunRepo,
		settingsRepo:    settingsRepo,
	}
	return srv
}

func (s *WorkflowService) InitSchedule(ctx context.Context) error {
	// 每日清理工作流运行历史
	app.GetScheduler().MustAdd("cleanupWorkflowHistoryRuns", "0 0 * * *", func() {
		s.cleanupHistoryRuns(context.Background())
	})

	// 初始化工作流调度器
	if err := s.dispatcher.Bootup(ctx); err != nil {
		panic(err)
	}

	// 注册工作流后台任务
	{
		workflows, err := s.workflowRepo.ListEnabledScheduled(ctx)
		if err != nil {
			return err
		}

		for _, workflow := range workflows {
			var errs []error

			err := app.GetScheduler().Add(fmt.Sprintf("workflow#%s", workflow.Id), workflow.TriggerCron, func() {
				_, err := s.StartRun(context.Background(), &dtos.WorkflowStartRunReq{
					WorkflowId: workflow.Id,
					RunTrigger: domain.WorkflowTriggerTypeScheduled,
				})
				if err != nil {
					app.GetLogger().Error(fmt.Sprintf("failed to start scheduled run for workflow #%s", workflow.Id), slog.Any("error", err))
				}
			})
			if err != nil {
				app.GetLogger().Error(fmt.Sprintf("failed to register cron job for workflow #%s", workflow.Id), slog.Any("error", err))
				errs = append(errs, err)
			} else {
				app.GetLogger().Info(fmt.Sprintf("registered cron job for workflow #%s", workflow.Id), slog.String("cron", workflow.TriggerCron))
			}

			if len(errs) > 0 {
				return errors.Join(errs...)
			}
		}
	}

	return nil
}

func (s *WorkflowService) GetStatistics(ctx context.Context) (*dtos.WorkflowStatisticsResp, error) {
	stats := s.dispatcher.GetStatistics()
	return &dtos.WorkflowStatisticsResp{
		Concurrency:      stats.Concurrency,
		PendingRunIds:    stats.PendingRunIds,
		ProcessingRunIds: stats.ProcessingRunIds,
	}, nil
}

func (s *WorkflowService) StartRun(ctx context.Context, req *dtos.WorkflowStartRunReq) (*dtos.WorkflowStartRunResp, error) {
	workflow, err := s.workflowRepo.GetById(ctx, req.WorkflowId)
	if err != nil {
		return nil, err
	}

	if req.RunTrigger == domain.WorkflowTriggerTypeManual && (workflow.LastRunStatus == domain.WorkflowRunStatusTypePending || workflow.LastRunStatus == domain.WorkflowRunStatusTypeProcessing) {
		return nil, errors.New("workflow is already pending or processing")
	} else if workflow.GraphContent == nil {
		return nil, errors.New("workflow graph content is empty")
	} else if err := workflow.GraphContent.Verify(); err != nil {
		return nil, fmt.Errorf("workflow graph content is invalid: %w", err)
	}

	workflowRun := &domain.WorkflowRun{
		WorkflowId: workflow.Id,
		Status:     domain.WorkflowRunStatusTypePending,
		Trigger:    req.RunTrigger,
		StartedAt:  time.Now(),
		Graph:      workflow.GraphContent.Clone(),
	}
	if resp, err := s.workflowRunRepo.Save(ctx, workflowRun); err != nil {
		return nil, err
	} else {
		workflowRun = resp
	}

	if err := s.dispatcher.Start(ctx, workflowRun.Id); err != nil {
		return nil, err
	}

	return &dtos.WorkflowStartRunResp{RunId: workflowRun.Id}, nil
}

func (s *WorkflowService) CancelRun(ctx context.Context, req *dtos.WorkflowCancelRunReq) (*dtos.WorkflowCancelRunResp, error) {
	workflow, err := s.workflowRepo.GetById(ctx, req.WorkflowId)
	if err != nil {
		return nil, err
	}

	workflowRun, err := s.workflowRunRepo.GetById(ctx, req.RunId)
	if err != nil {
		return nil, err
	} else if workflowRun.WorkflowId != workflow.Id {
		return nil, errors.New("workflow run not found")
	} else if workflowRun.Status != domain.WorkflowRunStatusTypePending && workflowRun.Status != domain.WorkflowRunStatusTypeProcessing {
		return nil, errors.New("workflow run is not pending or processing")
	}

	if err := s.dispatcher.Cancel(ctx, workflowRun.Id); err != nil {
		return nil, err
	}

	return &dtos.WorkflowCancelRunResp{}, nil
}

func (s *WorkflowService) Shutdown(ctx context.Context) {
	s.dispatcher.Shutdown(ctx)
}

func (s *WorkflowService) cleanupHistoryRuns(ctx context.Context) error {
	settings, err := s.settingsRepo.GetByName(ctx, "persistence")
	if err != nil {
		app.GetLogger().Error("failed to get persistence settings", slog.Any("error", err))
		return err
	}

	persistenceSettings := settings.Content.AsPersistence()
	if persistenceSettings.WorkflowRunsMaxDaysRetention != 0 {
		ret, err := s.workflowRunRepo.DeleteWhere(
			ctx,
			dbx.NewExp(fmt.Sprintf("status!='%s'", string(domain.WorkflowRunStatusTypePending))),
			dbx.NewExp(fmt.Sprintf("status!='%s'", string(domain.WorkflowRunStatusTypeProcessing))),
			dbx.NewExp(fmt.Sprintf("endedAt<DATETIME('now', '-%d days')", persistenceSettings.WorkflowRunsMaxDaysRetention)),
		)
		if err != nil {
			app.GetLogger().Error("failed to delete workflow history runs", slog.Any("error", err))
			return err
		}

		if ret > 0 {
			app.GetLogger().Info(fmt.Sprintf("cleanup %d workflow history runs", ret))
		}
	}

	return nil
}
