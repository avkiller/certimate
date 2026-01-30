import { getI18n, useTranslation } from "react-i18next";
import { Form, Input } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const BizDeployNodeConfigFieldsProviderBaishanCDN = () => {
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
        name={[parentNamePath, "domain"]}
        initialValue={initialValues.domain}
        label={t("workflow_node.deploy.form.baishan_cdn_domain.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.baishan_cdn_domain.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.baishan_cdn_domain.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "certificateId"]}
        initialValue={initialValues.certificateId}
        label={t("workflow_node.deploy.form.baishan_cdn_certificate_id.label")}
        extra={t("workflow_node.deploy.form.baishan_cdn_certificate_id.help")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.baishan_cdn_certificate_id.tooltip") }}></span>}
      >
        <Input allowClear type="number" placeholder={t("workflow_node.deploy.form.baishan_cdn_certificate_id.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    domain: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    domain: z.string().refine((v) => validDomainName(v, { allowWildcard: true }), t("common.errmsg.domain_invalid")),
    certificateId: z
      .union([z.string(), z.number().int()])
      .nullish()
      .refine((v) => {
        if (!v) return true;
        return /^\d+$/.test(v + "") && +v > 0;
      }, t("workflow_node.deploy.form.baishan_cdn_certificate_id.placeholder")),
  });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderBaishanCDN, {
  getInitialValues,
  getSchema,
});

export default _default;
