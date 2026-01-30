import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Select } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const BizDeployNodeConfigFieldsProviderAliyunWAF = () => {
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
        label={t("workflow_node.deploy.form.aliyun_waf_region.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_waf_region.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_waf_region.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "serviceVersion"]}
        initialValue={initialValues.serviceVersion}
        label={t("workflow_node.deploy.form.aliyun_waf_service_version.label")}
        rules={[formRule]}
      >
        <Select placeholder={t("workflow_node.deploy.form.aliyun_waf_service_version.placeholder")}>
          <Select.Option key="3.0" value="3.0">
            3.0
          </Select.Option>
        </Select>
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "instanceId"]}
        initialValue={initialValues.instanceId}
        label={t("workflow_node.deploy.form.aliyun_waf_instance_id.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_waf_instance_id.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_waf_instance_id.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "domain"]}
        initialValue={initialValues.domain}
        label={t("workflow_node.deploy.form.aliyun_waf_domain.label")}
        extra={t("workflow_node.deploy.form.aliyun_waf_domain.help")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_waf_domain.tooltip") }}></span>}
      >
        <Input allowClear placeholder={t("workflow_node.deploy.form.aliyun_waf_domain.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    region: "",
    serviceVersion: "3.0",
    instanceId: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    region: z.string().nonempty(t("workflow_node.deploy.form.aliyun_waf_region.placeholder")),
    serviceVersion: z.literal("3.0", t("workflow_node.deploy.form.aliyun_waf_service_version.placeholder")),
    instanceId: z.string().nonempty(t("workflow_node.deploy.form.aliyun_waf_instance_id.placeholder")),
    domain: z
      .string()
      .nullish()
      .refine((v) => {
        return !v || validDomainName(v!, { allowWildcard: true });
      }, t("common.errmsg.domain_invalid")),
  });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderAliyunWAF, {
  getInitialValues,
  getSchema,
});

export default _default;
