import { getI18n, useTranslation } from "react-i18next";
import { Form, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { useFormNestedFieldsContext } from "./_context";

const AccessConfigFormFieldsProviderWestcn = () => {
  const { i18n, t } = useTranslation();

  const { parentNamePath } = useFormNestedFieldsContext();
  const formSchema = z.object({
    [parentNamePath]: getSchema({ i18n }),
  });
  const formRule = createSchemaFieldRule(formSchema);
  const initialValues = getInitialValues();

  return (
    <>
      <Form.Item
        name={[parentNamePath, "username"]}
        initialValue={initialValues.username}
        label={t("access.form.westcn_username.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.westcn_username.tooltip") }}></span>}
      >
        <Input autoComplete="new-password" placeholder={t("access.form.westcn_username.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "apiPassword"]}
        initialValue={initialValues.apiPassword}
        label={t("access.form.westcn_api_password.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.westcn_api_password.tooltip") }}></span>}
      >
        <Input.Password autoComplete="new-password" placeholder={t("access.form.westcn_api_password.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    username: "",
    apiPassword: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    username: z
      .string()
      .min(1, t("access.form.westcn_username.placeholder"))
      .max(64, t("common.errmsg.string_max", { max: 64 })),
    apiPassword: z
      .string()
      .min(1, t("access.form.westcn_api_password.placeholder"))
      .max(64, t("common.errmsg.string_max", { max: 64 })),
  });
};

const _default = Object.assign(AccessConfigFormFieldsProviderWestcn, {
  getInitialValues,
  getSchema,
});

export default _default;
