import { getI18n } from "react-i18next";
import { z } from "zod";

const BizDeployNodeConfigFieldsProviderBaotaPanelConsoleGo = () => {
  return <></>;
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {};
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t: _ } = i18n;

  return z.object({});
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderBaotaPanelConsoleGo, {
  getInitialValues,
  getSchema,
});

export default _default;
