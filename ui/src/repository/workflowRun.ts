import { type RecordSubscription } from "pocketbase";

import { type WorkflowRunModel } from "@/domain/workflowRun";

import { COLLECTION_NAME_WORKFLOW_RUN, getPocketBase } from "./_pocketbase";

export type ListRequest = {
  workflowId?: string;
  page?: number;
  perPage?: number;
  expand?: boolean;
};

export const list = async (request: ListRequest) => {
  const pb = getPocketBase();

  const filters: string[] = [];
  if (request.workflowId) {
    filters.push(pb.filter("workflowRef={:workflowId}", { workflowId: request.workflowId }));
  }

  const page = request.page || 1;
  const perPage = request.perPage || 10;
  return await pb.collection(COLLECTION_NAME_WORKFLOW_RUN).getList<WorkflowRunModel>(page, perPage, {
    expand: request.expand ? ["workflowRef"].join(",") : void 0,
    fields: [
      "id",
      "status",
      "trigger",
      "startedAt",
      "endedAt",
      "error",
      "created",
      "updated",
      "deleted",
      "expand.workflowRef.id",
      "expand.workflowRef.name",
      "expand.workflowRef.description",
    ].join(","),
    filter: filters.join(" && "),
    sort: "-created",
    requestKey: null,
  });
};

export const get = async (id: string) => {
  return await getPocketBase()
    .collection(COLLECTION_NAME_WORKFLOW_RUN)
    .getOne<WorkflowRunModel>(id, {
      expand: ["workflowRef"].join(","),
      fields: ["*", "expand.workflowRef.id", "expand.workflowRef.name", "expand.workflowRef.description"].join(","),
      requestKey: null,
    });
};

export const remove = async (record: MaybeModelRecordWithId<WorkflowRunModel> | MaybeModelRecordWithId<WorkflowRunModel>[]) => {
  const pb = getPocketBase();

  if (Array.isArray(record)) {
    const batch = pb.createBatch();
    for (const item of record) {
      batch.collection(COLLECTION_NAME_WORKFLOW_RUN).delete(item.id);
    }
    const res = await batch.send();
    return res.every((e) => e.status >= 200 && e.status < 400);
  } else {
    await pb.collection(COLLECTION_NAME_WORKFLOW_RUN).delete(record.id!);
    return true;
  }
};

export const subscribe = async (id: string, cb: (e: RecordSubscription<WorkflowRunModel>) => void) => {
  return getPocketBase().collection(COLLECTION_NAME_WORKFLOW_RUN).subscribe(id, cb);
};

export const unsubscribe = async (id: string) => {
  return getPocketBase().collection(COLLECTION_NAME_WORKFLOW_RUN).unsubscribe(id);
};
