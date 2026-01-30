import dayjs from "dayjs";

import { type CertificateModel } from "@/domain/certificate";
import { COLLECTION_NAME_CERTIFICATE, getPocketBase } from "./_pocketbase";

const _commonFields = [
  "id",
  "source",
  "subjectAltNames",
  "serialNumber",
  "issuerOrg",
  "keyAlgorithm",
  "validityNotBefore",
  "validityNotAfter",
  "validityInterval",
  "isRenewed",
  "isRevoked",
  "workflowRef",
  "created",
  "updated",
  "deleted",
];
const _expandFields = ["expand.workflowRef.id", "expand.workflowRef.name", "expand.workflowRef.description"];

export const list = async ({
  keyword,
  state,
  stateThreshold,
  sort = "-created",
  page = 1,
  perPage = 10,
}: {
  keyword?: string;
  state?: "expiringSoon" | "expired";
  stateThreshold?: number;
  sort?: string;
  page?: number;
  perPage?: number;
}) => {
  const pb = getPocketBase();

  const filters: string[] = ["deleted=null"];
  if (keyword) {
    filters.push(pb.filter("(id={:keyword} || serialNumber={:keyword} || subjectAltNames~{:keyword})", { keyword: keyword }));
  }
  if (state === "expiringSoon") {
    filters.push(pb.filter("validityNotAfter<={:expiredAt}", { expiredAt: dayjs().add(stateThreshold!, "d").toDate() }));
    filters.push(pb.filter("validityNotAfter>@now"));
    filters.push(pb.filter("isRevoked=0"));
  } else if (state === "expired") {
    filters.push(pb.filter("validityNotAfter<=@now"));
  }

  return pb.collection(COLLECTION_NAME_CERTIFICATE).getList<CertificateModel>(page, perPage, {
    expand: ["workflowRef"].join(","),
    fields: [..._commonFields, ..._expandFields].join(","),
    filter: filters.join(" && "),
    sort: sort || "-created",
    requestKey: null,
  });
};

export const listByWorkflowRunId = async (workflowRunId: string) => {
  const pb = getPocketBase();

  const list = await pb.collection(COLLECTION_NAME_CERTIFICATE).getFullList<CertificateModel>({
    batch: 65535,
    fields: [..._commonFields, ..._expandFields, "certificate", "privateKey"].join(","),
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
      fields: ["*", ..._expandFields].join(","),
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
