import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Select } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import Show from "@/components/Show";

import { useFormNestedFieldsContext } from "./_context";

const RESOURCE_TYPE_LOADBALANCER = "loadbalancer" as const;
const RESOURCE_TYPE_LISTENER = "listener" as const;
const RESOURCE_TYPE_CERTIFICATE = "certificate" as const;

const BizDeployNodeConfigFieldsProviderHuaweiCloudELB = () => {
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
        label={t("workflow_node.deploy.form.huaweicloud_elb_region.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_elb_region.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.huaweicloud_elb_region.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "resourceType"]}
        initialValue={initialValues.resourceType}
        label={t("workflow_node.deploy.form.shared_resource_type.label")}
        rules={[formRule]}
      >
        <Select
          options={[RESOURCE_TYPE_LOADBALANCER, RESOURCE_TYPE_LISTENER, RESOURCE_TYPE_CERTIFICATE].map((s) => ({
            value: s,
            label: t(`workflow_node.deploy.form.huaweicloud_elb_resource_type.option.${s}.label`),
          }))}
          placeholder={t("workflow_node.deploy.form.shared_resource_type.placeholder")}
        />
      </Form.Item>

      <Show when={fieldResourceType === RESOURCE_TYPE_CERTIFICATE}>
        <Form.Item
          name={[parentNamePath, "certificateId"]}
          initialValue={initialValues.certificateId}
          label={t("workflow_node.deploy.form.huaweicloud_elb_certificate_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_elb_certificate_id.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.huaweicloud_elb_certificate_id.placeholder")} />
        </Form.Item>
      </Show>

      <Show when={fieldResourceType === RESOURCE_TYPE_LOADBALANCER}>
        <Form.Item
          name={[parentNamePath, "loadbalancerId"]}
          initialValue={initialValues.loadbalancerId}
          label={t("workflow_node.deploy.form.huaweicloud_elb_loadbalancer_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_elb_loadbalancer_id.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.huaweicloud_elb_loadbalancer_id.placeholder")} />
        </Form.Item>
      </Show>

      <Show when={fieldResourceType === RESOURCE_TYPE_LISTENER}>
        <Form.Item
          name={[parentNamePath, "listenerId"]}
          initialValue={initialValues.listenerId}
          label={t("workflow_node.deploy.form.huaweicloud_elb_listener_id.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.huaweicloud_elb_listener_id.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.huaweicloud_elb_listener_id.placeholder")} />
        </Form.Item>
      </Show>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    region: "",
    resourceType: RESOURCE_TYPE_LISTENER,
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z
    .object({
      region: z.string().nonempty(t("workflow_node.deploy.form.huaweicloud_elb_region.placeholder")),
      resourceType: z.literal(
        [RESOURCE_TYPE_LOADBALANCER, RESOURCE_TYPE_LISTENER, RESOURCE_TYPE_CERTIFICATE],
        t("workflow_node.deploy.form.shared_resource_type.placeholder")
      ),
      loadbalancerId: z.string().nullish(),
      listenerId: z.string().nullish(),
      certificateId: z.string().nullish(),
    })
    .superRefine((values, ctx) => {
      switch (values.resourceType) {
        case RESOURCE_TYPE_LOADBALANCER:
          {
            if (!values.loadbalancerId?.trim()) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.huaweicloud_elb_loadbalancer_id.placeholder"),
                path: ["loadbalancerId"],
              });
            }
          }
          break;

        case RESOURCE_TYPE_LISTENER:
          {
            if (!values.listenerId?.trim()) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.huaweicloud_elb_listener_id.placeholder"),
                path: ["listenerId"],
              });
            }
          }
          break;

        case RESOURCE_TYPE_CERTIFICATE:
          {
            if (!values.certificateId?.trim()) {
              ctx.addIssue({
                code: "custom",
                message: t("workflow_node.deploy.form.huaweicloud_elb_certificate_id.placeholder"),
                path: ["certificateId"],
              });
            }
          }
          break;
      }
    });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderHuaweiCloudELB, {
  getInitialValues,
  getSchema,
});

export default _default;
