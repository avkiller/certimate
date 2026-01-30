import { getI18n, useTranslation } from "react-i18next";
import { Form, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { useFormNestedFieldsContext } from "./_context";

const BizDeployNodeConfigFieldsProviderAliyunESA = () => {
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
        name={[parentNamePath, "region"]}
        initialValue={initialValues.region}
        label={t("workflow_node.deploy.form.aliyun_esa_region.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_esa_region.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_esa_region.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "siteId"]}
        initialValue={initialValues.siteId}
        label={t("workflow_node.deploy.form.aliyun_esa_site_id.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_esa_site_id.tooltip") }}></span>}
      >
        <Input type="number" placeholder={t("workflow_node.deploy.form.aliyun_esa_site_id.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    region: "",
    siteId: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    region: z.string().nonempty(t("workflow_node.deploy.form.aliyun_esa_region.placeholder")),
    siteId: z.union([z.string(), z.number().int()]).refine((v) => {
      return /^\d+$/.test(v + "") && +v > 0;
    }, t("workflow_node.deploy.form.aliyun_esa_site_id.placeholder")),
  });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderAliyunESA, {
  getInitialValues,
  getSchema,
});

export default _default;
