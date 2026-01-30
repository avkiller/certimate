import { getI18n, useTranslation } from "react-i18next";
import { Form, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import CodeInput from "@/components/CodeInput";

import { useFormNestedFieldsContext } from "./_context";

const BizNotifyNodeConfigFieldsProviderWebhook = () => {
  const { i18n, t } = useTranslation();

  const { parentNamePath } = useFormNestedFieldsContext();
  const formSchema = z.object({
    [parentNamePath]: getSchema({ i18n }),
  });
  const formRule = createSchemaFieldRule(formSchema);
  const formInst = Form.useFormInstance();
  const initialValues = getInitialValues();

  const handleWebhookDataBlur = () => {
    const value = formInst.getFieldValue([parentNamePath, "webhookData"]);
    try {
      const json = JSON.stringify(JSON.parse(value), null, 2);
      formInst.setFieldValue([parentNamePath, "webhookData"], json);
    } catch {
      return;
    }
  };

  return (
    <>
      <Form.Item
        name={[parentNamePath, "webhookData"]}
        initialValue={initialValues.webhookData}
        label={t("workflow_node.notify.form.webhook_data.label")}
        extra={t("workflow_node.notify.form.webhook_data.help")}
        rules={[formRule]}
      >
        <CodeInput
          height="auto"
          minHeight="64px"
          maxHeight="256px"
          language="json"
          placeholder={t("workflow_node.notify.form.webhook_data.placeholder")}
          onBlur={handleWebhookDataBlur}
        />
      </Form.Item>

      <Form.Item name={[parentNamePath, "timeout"]} label={t("workflow_node.notify.form.webhook_timeout.label")} rules={[formRule]}>
        <Input
          type="number"
          allowClear
          min={0}
          max={3600}
          placeholder={t("workflow_node.notify.form.webhook_timeout.placeholder")}
          addonAfter={t("workflow_node.notify.form.webhook_timeout.unit")}
        />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {};
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    webhookData: z
      .string()
      .nullish()
      .refine((v) => {
        if (!v) return true;

        try {
          const obj = JSON.parse(v);
          return typeof obj === "object" && !Array.isArray(obj);
        } catch {
          return false;
        }
      }, t("workflow_node.notify.form.webhook_data.errmsg.json_invalid")),
    timeout: z.preprocess(
      (v) => (v == null || v === "" ? void 0 : Number(v)),
      z.number().int(t("workflow_node.notify.form.webhook_timeout.placeholder")).gte(1, t("workflow_node.notify.form.webhook_timeout.placeholder")).nullish()
    ),
  });
};

const _default = Object.assign(BizNotifyNodeConfigFieldsProviderWebhook, {
  getInitialValues,
  getSchema,
});

export default _default;
