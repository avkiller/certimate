import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Select } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import Show from "@/components/Show";

import { useFormNestedFieldsContext } from "./_context";

const RESOURCE_TYPE_WEBSITE = "website" as const;
const RESOURCE_TYPE_CERTIFICATE = "certificate" as const;

const BizDeployNodeConfigFieldsProvider1PanelSite = () => {
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
        name={[parentNamePath, "nodeName"]}
        initialValue={initialValues.nodeName}
        label={t("workflow_node.deploy.form.1panel_site_node_name.label")}
        extra={t("workflow_node.deploy.form.1panel_site_node_name.help")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.1panel_site_node_name.tooltip") }}></span>}
      >
        <Input allowClear placeholder={t("workflow_node.deploy.form.1panel_site_node_name.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "resourceType"]}
        initialValue={initialValues.resourceType}
        label={t("workflow_node.deploy.form.1panel_site_resource_type.label")}
        rules={[formRule]}
      >
        <Select placeholder={t("workflow_node.deploy.form.1panel_site_resource_type.placeholder")}>
          <Select.Option key={RESOURCE_TYPE_WEBSITE} value={RESOURCE_TYPE_WEBSITE}>
            {t("workflow_node.deploy.form.1panel_site_resource_type.option.website.label")}
          </Select.Option>
          <Select.Option key={RESOURCE_TYPE_CERTIFICATE} value={RESOURCE_TYPE_CERTIFICATE}>
            {t("workflow_node.deploy.form.1panel_site_resource_type.option.certificate.label")}
          </Select.Option>
        </Select>
      </Form.Item>

      <Show when={fieldResourceType === RESOURCE_TYPE_WEBSITE}>
        <Form.Item
          name={[parentNamePath, "websiteId"]}
          initialValue={initialValues.websiteId}
          label={t("workflow_node.deploy.form.1panel_site_website_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.1panel_site_website_id.tooltip") }}></span>}
        >
          <Input type="number" placeholder={t("workflow_node.deploy.form.1panel_site_website_id.placeholder")} />
        </Form.Item>
      </Show>

      <Show when={fieldResourceType === RESOURCE_TYPE_CERTIFICATE}>
        <Form.Item
          name={[parentNamePath, "certificateId"]}
          initialValue={initialValues.certificateId}
          label={t("workflow_node.deploy.form.1panel_site_certificate_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.1panel_site_certificate_id.tooltip") }}></span>}
        >
          <Input type="number" placeholder={t("workflow_node.deploy.form.1panel_site_certificate_id.placeholder")} />
        </Form.Item>
      </Show>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    resourceType: RESOURCE_TYPE_WEBSITE,
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z
    .object({
      nodeName: z.string().nullish(),
      resourceType: z.literal([RESOURCE_TYPE_WEBSITE, RESOURCE_TYPE_CERTIFICATE], t("workflow_node.deploy.form.1panel_site_resource_type.placeholder")),
      websiteId: z.union([z.string(), z.number()]).nullish(),
      certificateId: z.union([z.string(), z.number()]).nullish(),
    })
    .superRefine((values, ctx) => {
      switch (values.resourceType) {
        case RESOURCE_TYPE_WEBSITE:
          {
            const res = z.preprocess((v) => Number(v), z.number().int().positive()).safeParse(values.websiteId);
            if (!res.success) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.1panel_site_website_id.placeholder"),
                path: ["websiteId"],
              });
            }
          }
          break;

        case RESOURCE_TYPE_CERTIFICATE:
          {
            const res = z.preprocess((v) => Number(v), z.number().int().positive()).safeParse(values.certificateId);
            if (!res.success) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.1panel_site_certificate_id.placeholder"),
                path: ["websiteId"],
              });
            }
          }
          break;
      }
    });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProvider1PanelSite, {
  getInitialValues,
  getSchema,
});

export default _default;
