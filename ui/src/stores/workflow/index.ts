import { create } from "zustand";

import {
  type WorkflowBranchNode,
  type WorkflowModel,
  type WorkflowNode,
  addBranch,
  addNode,
  getExecuteMethod,
  getWorkflowOutputBeforeId,
  removeBranch,
  removeNode,
  updateNode,
} from "@/domain/workflow";
import { get as getWorkflow, save as saveWorkflow } from "@/repository/workflow";

export type WorkflowState = {
  workflow: WorkflowModel;
  initialized: boolean;
  updateNode: (node: WorkflowNode) => void;
  addNode: (node: WorkflowNode, preId: string) => void;
  addBranch: (branchId: string) => void;
  removeNode: (nodeId: string) => void;
  removeBranch: (branchId: string, index: number) => void;
  getWorkflowOuptutBeforeId: (id: string, type: string) => WorkflowNode[];
  switchEnable(): void;
  save(): void;
  setBaseInfo: (name: string, description: string) => void;
  init(id: string): void;
  destroy(): void;
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

  setBaseInfo: async (name: string, description: string) => {
    const data: Record<string, string | boolean | WorkflowNode> = {
      id: (get().workflow.id as string) ?? "",
      name: name || "",
      description: description || "",
    };
    const resp = await saveWorkflow(data);
    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          name,
          description,
          id: resp.id,
        },
      };
    });
  },

  switchEnable: async () => {
    const root = get().workflow.content as WorkflowNode;
    const executeMethod = getExecuteMethod(root);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      content: root,
      enabled: !get().workflow.enabled,
      type: executeMethod.type,
      crontab: executeMethod.crontab,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          id: resp.id,
          content: resp.content,
          enabled: resp.enabled,
          type: resp.type,
          crontab: resp.crontab,
        },
      };
    });
  },

  save: async () => {
    const root = get().workflow.draft as WorkflowNode;
    const executeMethod = getExecuteMethod(root);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      content: root,
      hasDraft: false,
      type: executeMethod.type,
      crontab: executeMethod.crontab,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          id: resp.id,
          content: resp.content,
          hasDraft: false,
          type: resp.type,
          crontab: resp.crontab,
        },
      };
    });
  },

  updateNode: async (node: WorkflowNode | WorkflowBranchNode) => {
    const newRoot = updateNode(get().workflow.draft as WorkflowNode, node);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      draft: newRoot,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          draft: newRoot,
          id: resp.id,
          hasDraft: true,
        },
      };
    });
  },

  addNode: async (node: WorkflowNode | WorkflowBranchNode, preId: string) => {
    const newRoot = addNode(get().workflow.draft as WorkflowNode, preId, node);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      draft: newRoot,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          draft: newRoot,
          id: resp.id,
          hasDraft: true,
        },
      };
    });
  },

  addBranch: async (branchId: string) => {
    const newRoot = addBranch(get().workflow.draft as WorkflowNode, branchId);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      draft: newRoot,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          draft: newRoot,
          id: resp.id,
          hasDraft: true,
        },
      };
    });
  },

  removeBranch: async (branchId: string, index: number) => {
    const newRoot = removeBranch(get().workflow.draft as WorkflowNode, branchId, index);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      draft: newRoot,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          draft: newRoot,
          id: resp.id,
          hasDraft: true,
        },
      };
    });
  },

  removeNode: async (nodeId: string) => {
    const newRoot = removeNode(get().workflow.draft as WorkflowNode, nodeId);
    const resp = await saveWorkflow({
      id: (get().workflow.id as string) ?? "",
      draft: newRoot,
      hasDraft: true,
    });

    set((state: WorkflowState) => {
      return {
        workflow: {
          ...state.workflow,
          draft: newRoot,
          id: resp.id,
          hasDraft: true,
        },
      };
    });
  },

  getWorkflowOuptutBeforeId: (id: string, type: string) => {
    return getWorkflowOutputBeforeId(get().workflow.draft as WorkflowNode, id, type);
  },

  destroy: () => {
    set({
      workflow: {} as WorkflowModel,
      initialized: false,
    });
  },
}));
