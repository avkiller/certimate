import { getPocketBase } from "@/repository/pocketbase";

export const notifyTest = async (channel: string) => {
  const pb = getPocketBase();

  const resp = await pb.send("/api/notify/test", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: {
      channel,
    },
  });

  if (resp.code != 0) {
    throw new Error(resp.msg);
  }

  return resp;
};
