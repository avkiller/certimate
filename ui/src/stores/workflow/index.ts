import { produce } from "immer";
import { create } from "zustand";

import {
  type WorkflowModel,
  type WorkflowNode,
  type WorkflowNodeConfigForStart,
  addBranch,
  addNode,
  getWorkflowOutputBeforeId,
  removeBranch,
  removeNode,
  updateNode,
} from "@/domain/workflow";
import { get as getWorkflow, save as saveWorkflow } from "@/repository/workflow";

export type WorkflowState = {
  workflow: WorkflowModel;
  initialized: boolean;

  init(id: string): void;
  setBaseInfo: (name: string, description: string) => void;
  setEnabled(enabled: boolean): void;
  release(): void;
  discard(): void;
  destroy(): void;

  addNode: (node: WorkflowNode, previousNodeId: string) => void;
  updateNode: (node: WorkflowNode) => void;
  removeNode: (nodeId: string) => void;

  addBranch: (branchId: string) => void;
  removeBranch: (branchId: string, index: number) => void;

  getWorkflowOuptutBeforeId: (nodeId: string, type: string) => WorkflowNode[];
};

export const useWorkflowStore = create<WorkflowState>((set, get) => ({
  workflow: {} as WorkflowModel,
  initialized: false,

  init: async (id: string) => {
    const data = await getWorkflow(id);

    set({
      workflow: data,
      initialized: true,
    });
  },

  destroy: () => {
    set({
      workflow: {} as WorkflowModel,
      initialized: false,
    });
  },

  setBaseInfo: async (name: string, description: string) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const resp = await saveWorkflow({
      id: get().workflow.id!,
      name: name || "",
      description: description || "",
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.name = resp.name;
          draft.description = resp.description;
        }),
      };
    });
  },

  setEnabled: async (enabled: boolean) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const resp = await saveWorkflow({
      id: get().workflow.id!,
      enabled: enabled,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.enabled = resp.enabled;
        }),
      };
    });
  },

  release: async () => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = get().workflow.draft!;
    const startConfig = root.config as WorkflowNodeConfigForStart;
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      trigger: startConfig.trigger,
      triggerCron: startConfig.triggerCron,
      content: root,
      hasDraft: false,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.trigger = resp.trigger;
          draft.triggerCron = resp.triggerCron;
          draft.content = resp.content;
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  discard: async () => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = get().workflow.content!;
    const startConfig = root.config as WorkflowNodeConfigForStart;
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      draft: root,
      hasDraft: false,
      trigger: startConfig.trigger,
      triggerCron: startConfig.triggerCron,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.trigger = resp.trigger;
          draft.triggerCron = resp.triggerCron;
          draft.content = resp.content;
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  addNode: async (node: WorkflowNode, previousNodeId: string) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = addNode(get().workflow.draft!, previousNodeId, node);
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      draft: root,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  updateNode: async (node: WorkflowNode) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = updateNode(get().workflow.draft!, node);
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      draft: root,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  removeNode: async (nodeId: string) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = removeNode(get().workflow.draft!, nodeId);
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      draft: root,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  addBranch: async (branchId: string) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = addBranch(get().workflow.draft!, branchId);
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      draft: root,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  removeBranch: async (branchId: string, index: number) => {
    if (!get().initialized) throw "Workflow not initialized yet";

    const root = removeBranch(get().workflow.draft!, branchId, index);
    const resp = await saveWorkflow({
      id: get().workflow.id!,
      draft: root,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: produce(state.workflow, (draft) => {
          draft.draft = resp.draft;
          draft.hasDraft = resp.hasDraft;
        }),
      };
    });
  },

  getWorkflowOuptutBeforeId: (nodeId: string, type: string) => {
    return getWorkflowOutputBeforeId(get().workflow.draft as WorkflowNode, nodeId, type);
  },
}));
