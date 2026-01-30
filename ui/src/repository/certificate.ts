import dayjs from "dayjs";

import { type CertificateModel } from "@/domain/certificate";
import { COLLECTION_NAME_CERTIFICATE, getPocketBase } from "./_pocketbase";

export type ListRequest = {
  keyword?: string;
  state?: "expiringSoon" | "expired";
  sort?: string;
  page?: number;
  perPage?: number;
};

export const list = async (request: ListRequest) => {
  const pb = getPocketBase();

  const filters: string[] = ["deleted=null"];
  if (request.keyword) {
    filters.push(pb.filter("(id={:keyword} || serialNumber={:keyword} || subjectAltNames~{:keyword})", { keyword: request.keyword }));
  }
  if (request.state === "expiringSoon") {
    filters.push(pb.filter("validityNotAfter<{:expiredAt} && validityNotAfter>@now", { expiredAt: dayjs().add(20, "d").toDate() }));
  } else if (request.state === "expired") {
    filters.push(pb.filter("validityNotAfter<={:expiredAt}", { expiredAt: new Date() }));
  }

  const sort = request.sort || "-created";

  const page = request.page || 1;
  const perPage = request.perPage || 10;

  return pb.collection(COLLECTION_NAME_CERTIFICATE).getList<CertificateModel>(page, perPage, {
    expand: ["workflowRef"].join(","),
    fields: [
      "id",
      "source",
      "subjectAltNames",
      "serialNumber",
      "issuerOrg",
      "keyAlgorithm",
      "validityNotBefore",
      "validityNotAfter",
      "workflowRef",
      "created",
      "updated",
      "deleted",
      "expand.workflowRef.id",
      "expand.workflowRef.name",
      "expand.workflowRef.description",
    ].join(","),
    filter: filters.join(" && "),
    sort: sort,
    requestKey: null,
  });
};

export const listByWorkflowRunId = async (workflowRunId: string) => {
  const pb = getPocketBase();

  const list = await pb.collection(COLLECTION_NAME_CERTIFICATE).getFullList<CertificateModel>({
    batch: 65535,
    filter: pb.filter("workflowRunRef={:workflowRunId}", { workflowRunId }),
    sort: "created",
    requestKey: null,
  });

  return {
    totalItems: list.length,
    items: list,
  };
};

export const get = async (id: string) => {
  return await getPocketBase()
    .collection(COLLECTION_NAME_CERTIFICATE)
    .getOne<CertificateModel>(id, {
      expand: ["workflowRef"].join(","),
      fields: ["*", "expand.workflowRef.id", "expand.workflowRef.name", "expand.workflowRef.description"].join(","),
      requestKey: null,
    });
};

export const remove = async (record: MaybeModelRecordWithId<CertificateModel> | MaybeModelRecordWithId<CertificateModel>[]) => {
  const pb = getPocketBase();

  const deletedAt = dayjs.utc().format("YYYY-MM-DD HH:mm:ss");

  if (Array.isArray(record)) {
    const batch = pb.createBatch();
    for (const item of record) {
      batch.collection(COLLECTION_NAME_CERTIFICATE).update(item.id, { deleted: deletedAt });
    }
    const res = await batch.send();
    return res.every((e) => e.status >= 200 && e.status < 400);
  } else {
    await pb.collection(COLLECTION_NAME_CERTIFICATE).update<CertificateModel>(record.id!, { deleted: deletedAt });
    return true;
  }
};
