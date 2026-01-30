import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, InputNumber } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { validPortNumber } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const BizDeployNodeConfigFieldsProviderBaotaWAFSite = () => {
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
        name={[parentNamePath, "siteName"]}
        initialValue={initialValues.siteName}
        label={t("workflow_node.deploy.form.baotawaf_site_name.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.baotawaf_site_name.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.baotawaf_site_name.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "sitePort"]}
        initialValue={initialValues.sitePort}
        label={t("workflow_node.deploy.form.baotawaf_site_port.label")}
        rules={[formRule]}
      >
        <InputNumber style={{ width: "100%" }} placeholder={t("access.form.ssh_port.placeholder")} min={1} max={65535} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    siteName: "",
    sitePort: 443,
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    siteName: z.string().nonempty(t("workflow_node.deploy.form.baotawaf_site_name.placeholder")),
    sitePort: z.preprocess(
      (v) => Number(v),
      z
        .number()
        .int(t("workflow_node.deploy.form.baotawaf_site_port.placeholder"))
        .refine((v) => validPortNumber(v), t("common.errmsg.port_invalid"))
    ),
  });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderBaotaWAFSite, {
  getInitialValues,
  getSchema,
});

export default _default;
