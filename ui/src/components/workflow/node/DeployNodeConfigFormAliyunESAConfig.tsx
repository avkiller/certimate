import { useTranslation } from "react-i18next";
import { Form, type FormInstance, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

type DeployNodeConfigFormAliyunESAConfigFieldValues = Nullish<{
  region: string;
  siteId: string;
}>;

export type DeployNodeConfigFormAliyunESAConfigProps = {
  form: FormInstance;
  formName: string;
  disabled?: boolean;
  initialValues?: DeployNodeConfigFormAliyunESAConfigFieldValues;
  onValuesChange?: (values: DeployNodeConfigFormAliyunESAConfigFieldValues) => void;
};

const initFormModel = (): DeployNodeConfigFormAliyunESAConfigFieldValues => {
  return {};
};

const DeployNodeConfigFormAliyunESAConfig = ({
  form: formInst,
  formName,
  disabled,
  initialValues,
  onValuesChange,
}: DeployNodeConfigFormAliyunESAConfigProps) => {
  const { t } = useTranslation();

  const formSchema = z.object({
    region: z
      .string({ message: t("workflow_node.deploy.form.aliyun_esa_region.placeholder") })
      .nonempty(t("workflow_node.deploy.form.aliyun_esa_region.placeholder"))
      .trim(),
    siteId: z
      .string({ message: t("workflow_node.deploy.form.aliyun_esa_site_id.placeholder") })
      .regex(/^[1-9]\d*$/, t("workflow_node.deploy.form.aliyun_esa_site_id.placeholder"))
      .trim(),
  });
  const formRule = createSchemaFieldRule(formSchema);

  const handleFormChange = (_: unknown, values: z.infer<typeof formSchema>) => {
    onValuesChange?.(values);
  };

  return (
    <Form
      form={formInst}
      disabled={disabled}
      initialValues={initialValues ?? initFormModel()}
      layout="vertical"
      name={formName}
      onValuesChange={handleFormChange}
    >
      <Form.Item
        name="region"
        label={t("workflow_node.deploy.form.aliyun_esa_region.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_esa_region.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_esa_region.placeholder")} />
      </Form.Item>

      <Form.Item
        name="siteId"
        label={t("workflow_node.deploy.form.aliyun_esa_site_id.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.aliyun_esa_site_id.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.aliyun_esa_site_id.placeholder")} />
      </Form.Item>
    </Form>
  );
};

export default DeployNodeConfigFormAliyunESAConfig;
