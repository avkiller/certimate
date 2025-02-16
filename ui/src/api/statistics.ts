import { ClientResponseError } from "pocketbase";

import { type Statistics } from "@/domain/statistics";
import { getPocketBase } from "@/repository/_pocketbase";

export const get = async () => {
  const pb = getPocketBase();

  const resp = await pb.send<BaseResponse<Statistics>>("/api/statistics/get", {
    method: "GET",
  });

  if (resp.code != 0) {
    throw new ClientResponseError({ status: resp.code, response: resp, data: {} });
  }

  return resp;
};
