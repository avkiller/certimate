import { getI18n, useTranslation } from "react-i18next";
import { Form, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { useFormNestedFieldsContext } from "./_context";

const AccessConfigFormFieldsProviderDingTalkBot = () => {
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
        name={[parentNamePath, "webhookUrl"]}
        initialValue={initialValues.webhookUrl}
        label={t("access.form.dingtalkbot_webhook_url.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.dingtalkbot_webhook_url.tooltip") }}></span>}
      >
        <Input placeholder={t("access.form.dingtalkbot_webhook_url.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "secret"]}
        initialValue={initialValues.secret}
        label={t("access.form.dingtalkbot_secret.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("access.form.dingtalkbot_secret.tooltip") }}></span>}
      >
        <Input.Password autoComplete="new-password" placeholder={t("access.form.dingtalkbot_secret.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    webhookUrl: "",
    secret: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    webhookUrl: z.url(t("common.errmsg.url_invalid")),
    secret: z.string().nonempty(t("access.form.dingtalkbot_secret.placeholder")),
  });
};

const _default = Object.assign(AccessConfigFormFieldsProviderDingTalkBot, {
  getInitialValues,
  getSchema,
});

export default _default;
