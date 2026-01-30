import { getI18n, useTranslation } from "react-i18next";
import { Form, Input, InputNumber, Switch } from "antd";
import { createSchemaFieldRule } from "antd-zod";
import { z } from "zod";

import { validEmailAddress, validPortNumber } from "@/utils/validators";

import { useFormNestedFieldsContext } from "./_context";

const AccessConfigFormFieldsProviderEmail = () => {
  const { i18n, t } = useTranslation();

  const { parentNamePath } = useFormNestedFieldsContext();
  const formSchema = z.object({
    [parentNamePath]: getSchema({ i18n }),
  });
  const formRule = createSchemaFieldRule(formSchema);
  const formInst = Form.useFormInstance();
  const initialValues = getInitialValues();

  const handleTlsSwitchChange = (checked: boolean) => {
    const oldPort = formInst.getFieldValue([parentNamePath, "smtpPort"]);
    const newPort = checked && (oldPort == null || oldPort === 25) ? 465 : !checked && (oldPort == null || oldPort === 465) ? 25 : oldPort;
    if (newPort !== oldPort) {
      formInst.setFieldValue([parentNamePath, "smtpPort"], newPort);
    }
  };

  return (
    <>
      <div className="flex space-x-2">
        <div className="w-3/5">
          <Form.Item
            name={[parentNamePath, "smtpHost"]}
            initialValue={initialValues.smtpHost}
            label={t("access.form.email_smtp_host.label")}
            rules={[formRule]}
          >
            <Input placeholder={t("access.form.email_smtp_host.placeholder")} />
          </Form.Item>
        </div>

        <div className="w-2/5">
          <Form.Item
            name={[parentNamePath, "smtpPort"]}
            initialValue={initialValues.smtpPort}
            label={t("access.form.email_smtp_port.label")}
            rules={[formRule]}
          >
            <InputNumber style={{ width: "100%" }} placeholder={t("access.form.email_smtp_port.placeholder")} min={1} max={65535} />
          </Form.Item>
        </div>
      </div>

      <Form.Item name={[parentNamePath, "smtpTls"]} initialValue={initialValues.smtpTls} label={t("access.form.email_smtp_tls.label")} rules={[formRule]}>
        <Switch onChange={handleTlsSwitchChange} />
      </Form.Item>

      <Form.Item name={[parentNamePath, "username"]} initialValue={initialValues.username} label={t("access.form.email_username.label")} rules={[formRule]}>
        <Input autoComplete="new-password" placeholder={t("access.form.email_username.placeholder")} />
      </Form.Item>

      <Form.Item name={[parentNamePath, "password"]} initialValue={initialValues.password} label={t("access.form.email_password.label")} rules={[formRule]}>
        <Input.Password autoComplete="new-password" placeholder={t("access.form.email_password.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "senderAddress"]}
        initialValue={initialValues.senderAddress}
        label={t("access.form.email_sender_address.label")}
        rules={[formRule]}
      >
        <Input type="email" allowClear placeholder={t("access.form.email_sender_address.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "senderName"]}
        initialValue={initialValues.senderName}
        label={t("access.form.email_sender_name.label")}
        rules={[formRule]}
      >
        <Input allowClear placeholder={t("access.form.email_sender_name.placeholder")} />
      </Form.Item>

      <Form.Item
        name={[parentNamePath, "receiverAddress"]}
        initialValue={initialValues.receiverAddress}
        label={t("access.form.email_receiver_address.label")}
        extra={t("access.form.email_receiver_address.help")}
        rules={[formRule]}
      >
        <Input type="email" allowClear placeholder={t("access.form.email_receiver_address.placeholder")} />
      </Form.Item>
    </>
  );
};

const getInitialValues = (): Nullish<z.infer<ReturnType<typeof getSchema>>> => {
  return {
    smtpHost: "",
    smtpPort: 465,
    smtpTls: true,
    username: "",
    password: "",
    senderAddress: "",
    senderName: "",
    receiverAddress: "",
  };
};

const getSchema = ({ i18n = getI18n() }: { i18n: ReturnType<typeof getI18n> }) => {
  const { t } = i18n;

  return z.object({
    smtpHost: z
      .string()
      .min(1, t("access.form.email_smtp_host.placeholder"))
      .max(256, t("common.errmsg.string_max", { max: 256 })),
    smtpPort: z.preprocess(
      (v) => Number(v),
      z.number().refine((v) => validPortNumber(v), t("common.errmsg.port_invalid"))
    ),
    smtpTls: z.boolean().nullish(),
    username: z
      .string()
      .min(1, t("access.form.email_username.placeholder"))
      .max(256, t("common.errmsg.string_max", { max: 256 })),
    password: z
      .string()
      .min(1, t("access.form.email_password.placeholder"))
      .max(256, t("common.errmsg.string_max", { max: 256 })),
    senderAddress: z.email(t("common.errmsg.email_invalid")),
    senderName: z.string().nullish(),
    receiverAddress: z
      .string()
      .nullish()
      .refine((v) => {
        if (!v) return true;
        return validEmailAddress(v);
      }, t("common.errmsg.email_invalid")),
  });
};

const _default = Object.assign(AccessConfigFormFieldsProviderEmail, {
  getInitialValues,
  getSchema,
});

export default _default;
