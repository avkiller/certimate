import { ClientResponseError } from "pocketbase";

import { getPocketBase } from "@/repository/_pocketbase";

export const notifyTest = async (provider: string) => {
  const pb = getPocketBase();

  const resp = await pb.send<BaseResponse>("/api/notify/test", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: {
      provider,
    },
  });

  if (resp.code != 0) {
    throw new ClientResponseError({ status: resp.code, response: resp, data: {} });
  }

  return resp;
};
