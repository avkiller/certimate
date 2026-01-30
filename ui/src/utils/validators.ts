import { z } from "zod/mini";

import { validCronExpression as _validCronExpression } from "./cron";

export const validCronExpression = (value: string) => {
  return _validCronExpression(value);
};

export const validDomainName = (value: string, { allowWildcard = false }: { allowWildcard?: boolean } = {}) => {
  const re = allowWildcard
    ? /^(?:\*\.)?(?!-)[A-Za-z0-9-]{1,}(?<!-)(\.[A-Za-z0-9-]{1,}(?<!-)){0,}$/
    : /^(?!-)[A-Za-z0-9-]{1,}(?<!-)(\.[A-Za-z0-9-]{1,}(?<!-)){0,}$/;
  return re.test(value);
};

export const validEmailAddress = (value: string) => {
  return z.email().safeParse(value).success;
};

export const validIPv4Address = (value: string) => {
  return z.ipv4().safeParse(value).success;
};

export const validIPv6Address = (value: string) => {
  return z.ipv6().safeParse(value).success;
};

export const validHttpOrHttpsUrl = (value: string) => {
  try {
    const url = new URL(value);
    return url.protocol === "http:" || url.protocol === "https:";
  } catch {
    return false;
  }
};

export const validPortNumber = (value: string | number) => {
  return parseInt(value + "") === +value && String(+value) === String(value) && +value >= 1 && +value <= 65535;
};
