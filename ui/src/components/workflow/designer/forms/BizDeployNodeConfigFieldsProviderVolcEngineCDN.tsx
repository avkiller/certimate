import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Radio } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import Show from "@/components/Show";
import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const MATCH_PATTERN_EXACT = "exact" as const;
const MATCH_PATTERN_WILDCARD = "wildcard" as const;
const MATCH_PATTERN_CERTSAN = "certsan" as const;

const BizDeployNodeConfigFieldsProviderVolcEngineCDN = () => {
  const { i18n, t } = useTranslation();

  const { parentNamePath } = useFormNestedFieldsContext();
  const formSchema = z.object({
    [parentNamePath]: getSchema({ i18n }),
  });
  const formRule = createSchemaFieldRule(formSchema);
  const formInst = Form.useFormInstance();
  const initialValues = getInitialValues();

  const fieldDomainPattern = Form.useWatch([parentNamePath, "matchPattern"], { form: formInst, preserve: true });

  return (
    <>
      <Form.Item
        name={[parentNamePath, "matchPattern"]}
        initialValue={initialValues.matchPattern}
        label={t("workflow_node.deploy.form.shared_domain_match_pattern.label")}
        extra={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.shared_domain_match_pattern.help_wildcard") }}></span>}
        rules={[formRule]}
      >
        <Radio.Group
          options={[MATCH_PATTERN_EXACT, MATCH_PATTERN_WILDCARD, MATCH_PATTERN_CERTSAN].map((s) => ({
            key: s,
            label: t(`workflow_node.deploy.form.shared_domain_match_pattern.option.${s}.label`),
            value: s,
          }))}
        />
      </Form.Item>

      <Show when={fieldDomainPattern !== MATCH_PATTERN_CERTSAN}>
        <Form.Item
          name={[parentNamePath, "domain"]}
          initialValue={initialValues.domain}
          label={t("workflow_node.deploy.form.volcengine_cdn_domain.label")}
          rules={[formRule]}
          tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.volcengine_cdn_domain.tooltip") }}></span>}
        >
          <Input placeholder={t("workflow_node.deploy.form.volcengine_cdn_domain.placeholder")} />
        </Form.Item>
      </Show>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    matchPattern: MATCH_PATTERN_EXACT,
    domain: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n?: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z
    .object({
      matchPattern: z.enum(
        [MATCH_PATTERN_EXACT, MATCH_PATTERN_WILDCARD, MATCH_PATTERN_CERTSAN],
        t("workflow_node.deploy.form.shared_domain_match_pattern.placeholder")
      ),
      domain: z.string().nullish(),
    })
    .superRefine((values, ctx) => {
      if (values.matchPattern) {
        switch (values.matchPattern) {
          case MATCH_PATTERN_EXACT:
          case MATCH_PATTERN_WILDCARD:
            {
              if (!validDomainName(values.domain!, { allowWildcard: true })) {
                ctx.addIssue({
                  code: "custom",
                  message: t("common.errmsg.domain_invalid"),
                  path: ["domain"],
                });
              }
            }
            break;
        }
      }
    });
};

const _default = Object.assign(BizDeployNodeConfigFieldsProviderVolcEngineCDN, {
  getInitialValues,
  getSchema,
});

export default _default;
