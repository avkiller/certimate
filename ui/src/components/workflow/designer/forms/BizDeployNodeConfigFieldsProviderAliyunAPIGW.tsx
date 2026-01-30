import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Select } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import Show from "@/components/Show";
import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const SERVICE_TYPE_CLOUDNATIVE = "cloudnative" as const;
const SERVICE_TYPE_TRADITIONAL = "traditional" as const;

const BizDeployNodeConfigFieldsProviderAliyunAPIGW = () => {
  const { i18n, t } = useTranslation();

  const { parentNamePath } = useFormNestedFieldsContext();
  const formSchema = z.object({
    [parentNamePath]: getSchema({ i18n }),
  });
  const formRule = createSchemaFieldRule(formSchema);
  const formInst = Form.useFormInstance();
  const initialValues = getInitialValues();

  const fieldServiceType = Form.useWatch([parentNamePath, "serviceType"], formInst);

  return (
    <>
      <Form.Item
        name={[parentNamePath, "serviceType"]}
        initialValue={initialValues.serviceType}
        label={t("workflow_node.deploy.form.aliyun_apigw_service_type.label")}
        rules={[formRule]}
      >
        <Select placeholder={t("workflow_node.deploy.form.aliyun_apigw_service_type.placeholder")}>
          <Select.Option key={SERVICE_TYPE_CLOUDNATIVE} value={SERVICE_TYPE_CLOUDNATIVE}>
            {t("workflow_node.deploy.form.aliyun_apigw_service_type.option.cloudnative.label")}
          </Select.Option>
          <Select.Option key={SERVICE_TYPE_TRADITIONAL} value={SERVICE_TYPE_TRADITIONAL}>
            {t("workflow_node.deploy.form.aliyun_apigw_service_type.option.traditional.label")}
          </Select.Option>
        </Select>
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "region"]}
        initialValue={initialValues.region}
        label={t("workflow_node.deploy.form.aliyun_apigw_region.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_apigw_region.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_apigw_region.placeholder")} />
      </Form.Item>

      <Show when={fieldServiceType === SERVICE_TYPE_CLOUDNATIVE}>
        <Form.Item
          name={[parentNamePath, "gatewayId"]}
          initialValue={initialValues.gatewayId}
          label={t("workflow_node.deploy.form.aliyun_apigw_gateway_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_apigw_gateway_id.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.aliyun_apigw_gateway_id.placeholder")} />
        </Form.Item>
      </Show>

      <Show when={fieldServiceType === SERVICE_TYPE_TRADITIONAL}>
        <Form.Item
          name={[parentNamePath, "groupId"]}
          initialValue={initialValues.groupId}
          label={t("workflow_node.deploy.form.aliyun_apigw_group_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_apigw_group_id.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.aliyun_apigw_group_id.placeholder")} />
        </Form.Item>
      </Show>

      <Form.Item
        name={[parentNamePath, "domain"]}
        initialValue={initialValues.domain}
        label={t("workflow_node.deploy.form.aliyun_apigw_domain.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_apigw_domain.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_apigw_domain.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    region: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z
    .object({
      serviceType: z.literal([SERVICE_TYPE_CLOUDNATIVE, SERVICE_TYPE_TRADITIONAL], t("workflow_node.deploy.form.aliyun_apigw_service_type.placeholder")),
      region: z.string().nonempty(t("workflow_node.deploy.form.aliyun_apigw_region.placeholder")),
      gatewayId: z.string().nullish(),
      groupId: z.string().nullish(),
      domain: z
        .string()
        .nonempty(t("workflow_node.deploy.form.aliyun_apigw_domain.placeholder"))
        .refine((v) => validDomainName(v!, { allowWildcard: true }), t("common.errmsg.domain_invalid")),
    })
    .superRefine((values, ctx) => {
      switch (values.serviceType) {
        case SERVICE_TYPE_CLOUDNATIVE:
          {
            if (!values.gatewayId?.trim()) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.aliyun_apigw_gateway_id.placeholder"),
                path: ["gatewayId"],
              });
            }
          }
          break;

        case SERVICE_TYPE_TRADITIONAL:
          {
            if (!values.groupId?.trim()) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.aliyun_apigw_group_id.placeholder"),
                path: ["groupId"],
              });
            }
          }
          break;
      }
    });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderAliyunAPIGW, {
  getInitialValues,
  getSchema,
});

export default _default;
