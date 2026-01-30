import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, Radio } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { validDomainName } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const MATCH_PATTERN_EXACT = "exact" as const;
const MATCH_PATTERN_WILDCARD = "wildcard" as const;

const BizDeployNodeConfigFieldsProviderVolcEngineLive = () => {
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
        name={[parentNamePath, "matchPattern"]}
        initialValue={initialValues.matchPattern}
        label={t("workflow_node.deploy.form.shared_domain_match_pattern.label")}
        rules={[formRule]}
      >
        <Radio.Group
          options={[MATCH_PATTERN_EXACT, MATCH_PATTERN_WILDCARD].map((s) => ({
            key: s,
            label: t(`workflow_node.deploy.form.shared_domain_match_pattern.option.${s}.label`),
            value: s,
          }))}
        />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "domain"]}
        initialValue={initialValues.domain}
        label={t("workflow_node.deploy.form.volcengine_live_domain.label")}
        rules={[formRule]}
        tooltip={<span dangerouslySetInnerHTML={{ __html: t("workflow_node.deploy.form.volcengine_live_domain.tooltip") }}></span>}
      >
        <Input placeholder={t("workflow_node.deploy.form.volcengine_live_domain.placeholder")} />
      </Form.Item>
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
      matchPattern: z.enum([MATCH_PATTERN_EXACT, MATCH_PATTERN_WILDCARD], t("workflow_node.deploy.form.shared_domain_match_pattern.placeholder")),
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

const _default = Object.assign(BizDeployNodeConfigFieldsProviderVolcEngineLive, {
  getInitialValues,
  getSchema,
});

export default _default;
