import { getI18n, useTranslation } from "react-i18next";
import { Form } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import MultipleSplitValueInput from "@/components/MultipleSplitValueInput";
import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const MULTIPLE_INPUT_SEPARATOR = ";";

const BizDeployNodeConfigFieldsProviderWangsuCDN = () => {
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
        name={[parentNamePath, "domains"]}
        initialValue={initialValues.domains}
        label={t("workflow_node.deploy.form.wangsu_cdn_domains.label")}
        extra={t("workflow_node.deploy.form.wangsu_cdn_domains.help")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.wangsu_cdn_domains.tooltip") }}></span>}
      >
        <MultipleSplitValueInput
          modalTitle={t("workflow_node.deploy.form.wangsu_cdn_domains.multiple_input_modal.title")}
          placeholder={t("workflow_node.deploy.form.wangsu_cdn_domains.placeholder")}
          placeholderInModal={t("workflow_node.deploy.form.wangsu_cdn_domains.multiple_input_modal.placeholder")}
          separator={MULTIPLE_INPUT_SEPARATOR}
          splitOptions={{ removeEmpty: true, trimSpace: true }}
        />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    domains: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    domains: z.string().refine((v) => {
      if (!v) return false;
      return String(v)
        .split(MULTIPLE_INPUT_SEPARATOR)
        .every((e) => validDomainName(e, { allowWildcard: true }));
    }, t("workflow_node.deploy.form.wangsu_cdn_domains.placeholder")),
  });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderWangsuCDN, {
  getInitialValues,
  getSchema,
});

export default _default;
