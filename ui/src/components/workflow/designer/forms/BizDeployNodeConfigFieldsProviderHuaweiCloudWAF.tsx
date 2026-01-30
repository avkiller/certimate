import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Select } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import Show from "@/components/Show";
import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const RESOURCE_TYPE_CERTIFICATE = "certificate" as const;
const RESOURCE_TYPE_CLOUDSERVER = "cloudserver" as const;
const RESOURCE_TYPE_PREMIUMHOST = "premiumhost" as const;

const BizDeployNodeConfigFieldsProviderHuaweiCloudWAF = () => {
  const { i18n, t } = useTranslation();

  const { parentNamePath } = useFormNestedFieldsContext();
  const formSchema = z.object({
    [parentNamePath]: getSchema({ i18n }),
  });
  const formRule = createSchemaFieldRule(formSchema);
  const formInst = Form.useFormInstance();
  const initialValues = getInitialValues();

  const fieldResourceType = Form.useWatch([parentNamePath, "resourceType"], formInst);

  return (
    <>
      <Form.Item
        name={[parentNamePath, "region"]}
        initialValue={initialValues.region}
        label={t("workflow_node.deploy.form.huaweicloud_waf_region.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_waf_region.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.huaweicloud_waf_region.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "resourceType"]}
        initialValue={initialValues.resourceType}
        label={t("workflow_node.deploy.form.huaweicloud_waf_resource_type.label")}
        rules={[formRule]}
      >
        <Select placeholder={t("workflow_node.deploy.form.huaweicloud_waf_resource_type.placeholder")}>
          <Select.Option key={RESOURCE_TYPE_CERTIFICATE} value={RESOURCE_TYPE_CERTIFICATE}>
            {t("workflow_node.deploy.form.huaweicloud_waf_resource_type.option.certificate.label")}
          </Select.Option>
          <Select.Option key={RESOURCE_TYPE_CLOUDSERVER} value={RESOURCE_TYPE_CLOUDSERVER}>
            {t("workflow_node.deploy.form.huaweicloud_waf_resource_type.option.cloudserver.label")}
          </Select.Option>
          <Select.Option key={RESOURCE_TYPE_PREMIUMHOST} value={RESOURCE_TYPE_PREMIUMHOST}>
            {t("workflow_node.deploy.form.huaweicloud_waf_resource_type.option.premiumhost.label")}
          </Select.Option>
        </Select>
      </Form.Item>

      <Show when={fieldResourceType === RESOURCE_TYPE_CERTIFICATE}>
        <Form.Item
          name={[parentNamePath, "certificateId"]}
          initialValue={initialValues.certificateId}
          label={t("workflow_node.deploy.form.huaweicloud_waf_certificate_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_waf_certificate_id.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.huaweicloud_waf_certificate_id.placeholder")} />
        </Form.Item>
      </Show>

      <Show when={fieldResourceType === RESOURCE_TYPE_CLOUDSERVER || fieldResourceType === RESOURCE_TYPE_PREMIUMHOST}>
        <Form.Item
          name={[parentNamePath, "domain"]}
          initialValue={initialValues.domain}
          label={t("workflow_node.deploy.form.huaweicloud_waf_domain.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_waf_domain.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.huaweicloud_waf_domain.placeholder")} />
        </Form.Item>
      </Show>
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
      region: z.string().nonempty(t("workflow_node.deploy.form.huaweicloud_waf_region.placeholder")),
      resourceType: z.literal(
        [RESOURCE_TYPE_CERTIFICATE, RESOURCE_TYPE_CLOUDSERVER, RESOURCE_TYPE_PREMIUMHOST],
        t("workflow_node.deploy.form.huaweicloud_waf_resource_type.placeholder")
      ),
      certificateId: z.string().nullish(),
      domain: z.string().nullish(),
    })
    .superRefine((values, ctx) => {
      switch (values.resourceType) {
        case RESOURCE_TYPE_CERTIFICATE:
          {
            if (!values.certificateId?.trim()) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.huaweicloud_waf_certificate_id.placeholder"),
                path: ["certificateId"],
              });
            }
          }
          break;

        case RESOURCE_TYPE_CLOUDSERVER:
        case RESOURCE_TYPE_PREMIUMHOST:
          {
            if (!validDomainName(values.domain!, { allowWildcard: true })) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.huaweicloud_waf_domain.placeholder"),
                path: ["domain"],
              });
            }
          }
          break;
      }
    });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderHuaweiCloudWAF, {
  getInitialValues,
  getSchema,
});

export default _default;
