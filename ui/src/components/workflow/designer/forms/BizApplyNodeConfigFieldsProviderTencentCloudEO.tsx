import { getI18n, useTranslation } from "react-i18next";
import { Form, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { useFormNestedFieldsContext } from "./_context";

const BizApplyNodeConfigFieldsProviderTencentCloudEO = () => {
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
        name={[parentNamePath, "zoneId"]}
        initialValue={initialValues.zoneId}
        label={t("workflow_node.apply.form.tencentcloud_eo_zone_id.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.apply.form.tencentcloud_eo_zone_id.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.apply.form.tencentcloud_eo_zone_id.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    zoneId: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    zoneId: z.string().nonempty(t("workflow_node.apply.form.tencentcloud_eo_zone_id.placeholder")),
  });
};

const _default = Object.assign(BizApplyNodeConfigFieldsProviderTencentCloudEO, {
  getInitialValues,
  getSchema,
});

export default _default;
