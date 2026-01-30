import { ClientResponseError } from "pocketbase";

import { getPocketBase } from "@/repository/_pocketbase";

export const testPushNotification = async ({ provider, accessId }: { provider: string; accessId: string }) => {
  const pb = getPocketBase();

  const resp = await pb.send<BaseResponse>("/api/notifications/test", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: {
      provider,
      accessId,
    },
  });

  if (resp.code != 0) {
    throw new ClientResponseError({ status: resp.code, response: resp, data: {} });
  }

  return resp;
};
