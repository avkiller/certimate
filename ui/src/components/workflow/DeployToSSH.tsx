import { useTranslation } from "react-i18next";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "../ui/form";
import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from "../ui/select";
import { Button } from "../ui/button";
import { DeployFormProps } from "./DeployForm";
import { useWorkflowStore, WorkflowState } from "@/stores/workflow";
import { useShallow } from "zustand/shallow";
import { usePanel } from "./PanelProvider";
import { useEffect, useState } from "react";
import i18n from "@/i18n";
import { WorkflowNode } from "@/domain/workflow";
import { Textarea } from "../ui/textarea";
import AccessSelect from "./AccessSelect";
import AccessEditDialog from "../certimate/AccessEditDialog";
import { Plus } from "lucide-react";

const selectState = (state: WorkflowState) => ({
  updateNode: state.updateNode,
  getWorkflowOuptutBeforeId: state.getWorkflowOuptutBeforeId,
});

const t = i18n.t;

const formSchema = z
  .object({
    providerType: z.string(),
    access: z.string().min(1, t("domain.deployment.form.access.placeholder")),
    certificate: z.string().min(1),
    format: z.union([z.literal("pem"), z.literal("pfx"), z.literal("jks")], {
      message: t("domain.deployment.form.file_format.placeholder"),
    }),
    certPath: z
      .string()
      .min(1, t("domain.deployment.form.file_cert_path.placeholder"))
      .max(255, t("common.errmsg.string_max", { max: 255 })),
    keyPath: z
      .string()
      .min(0, t("domain.deployment.form.file_key_path.placeholder"))
      .max(255, t("common.errmsg.string_max", { max: 255 })),
    pfxPassword: z.string().optional(),
    jksAlias: z.string().optional(),
    jksKeypass: z.string().optional(),
    jksStorepass: z.string().optional(),
    preCommand: z.string().optional(),
    command: z.string().optional(),
  })
  .refine((data) => (data.format === "pem" ? !!data.keyPath?.trim() : true), {
    message: t("domain.deployment.form.file_key_path.placeholder"),
    path: ["keyPath"],
  })
  .refine((data) => (data.format === "pfx" ? !!data.pfxPassword?.trim() : true), {
    message: t("domain.deployment.form.file_pfx_password.placeholder"),
    path: ["pfxPassword"],
  })
  .refine((data) => (data.format === "jks" ? !!data.jksAlias?.trim() : true), {
    message: t("domain.deployment.form.file_jks_alias.placeholder"),
    path: ["jksAlias"],
  })
  .refine((data) => (data.format === "jks" ? !!data.jksKeypass?.trim() : true), {
    message: t("domain.deployment.form.file_jks_keypass.placeholder"),
    path: ["jksKeypass"],
  })
  .refine((data) => (data.format === "jks" ? !!data.jksStorepass?.trim() : true), {
    message: t("domain.deployment.form.file_jks_storepass.placeholder"),
    path: ["jksStorepass"],
  });

const DeployToSSH = ({ data }: DeployFormProps) => {
  const { updateNode, getWorkflowOuptutBeforeId } = useWorkflowStore(useShallow(selectState));
  const { hidePanel } = usePanel();
  const { t } = useTranslation();

  const [beforeOutput, setBeforeOutput] = useState<WorkflowNode[]>([]);

  useEffect(() => {
    const rs = getWorkflowOuptutBeforeId(data.id, "certificate");
    console.log(rs);
    setBeforeOutput(rs);
  }, [data]);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      providerType: "ssh",
      access: data.config?.access as string,
      certificate: data.config?.certificate as string,
      format: (data.config?.format as "pem" | "pfx" | "jks") || "pem",
      certPath: (data.config?.certPath as string) || "/etc/ssl/certs/cert.crt",
      keyPath: (data.config?.keyPath as string) || "/etc/ssl/private/cert.key",
      pfxPassword: (data.config?.pfxPassword as string) || "",
      jksAlias: (data.config?.jksAlias as string) || "",
      jksKeypass: (data.config?.jksKeypass as string) || "",
      jksStorepass: (data.config?.jksStorepass as string) || "",
      preCommand: (data.config?.preCommand as string) || "",
      command: (data.config?.command as string) || "service nginx reload",
    },
  });

  const format = form.watch("format");
  const certPath = form.watch("certPath");

  useEffect(() => {
    if (format === "pem" && /(.pfx|.jks)$/.test(certPath)) {
      form.setValue("certPath", certPath.replace(/(.pfx|.jks)$/, ".crt"));
    } else if (format === "pfx" && /(.crt|.jks)$/.test(certPath)) {
      form.setValue("certPath", certPath.replace(/(.crt|.jks)$/, ".pfx"));
    } else if (format === "jks" && /(.crt|.pfx)$/.test(certPath)) {
      form.setValue("certPath", certPath.replace(/(.crt|.pfx)$/, ".jks"));
    }
  }, [format]);

  const onSubmit = async (config: z.infer<typeof formSchema>) => {
    updateNode({ ...data, config, validated: true });
    hidePanel();
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="access"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="flex justify-between">
                <div>{t("domain.deployment.form.access.label")}</div>

                <AccessEditDialog
                  trigger={
                    <div className="font-normal text-primary hover:underline cursor-pointer flex items-center">
                      <Plus size={14} />
                      {t("common.button.add")}
                    </div>
                  }
                  op="add"
                  outConfigType="ssh"
                />
              </FormLabel>
              <FormControl>
                <AccessSelect
                  {...field}
                  value={field.value}
                  onValueChange={(value) => {
                    form.setValue("access", value);
                  }}
                  providerType="ssh"
                />
              </FormControl>

              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="certificate"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("workflow.common.certificate.label")}</FormLabel>
              <FormControl>
                <Select
                  {...field}
                  value={field.value}
                  onValueChange={(value) => {
                    form.setValue("certificate", value);
                  }}
                >
                  <SelectTrigger>
                    <SelectValue placeholder={t("workflow.common.certificate.placeholder")} />
                  </SelectTrigger>
                  <SelectContent>
                    {beforeOutput.map((item) => (
                      <>
                        <SelectGroup key={item.id}>
                          <SelectLabel>{item.name}</SelectLabel>
                          {item.output?.map((output) => (
                            <SelectItem key={output.name} value={`${item.id}#${output.name}`}>
                              <div>
                                {item.name}-{output.label}
                              </div>
                            </SelectItem>
                          ))}
                        </SelectGroup>
                      </>
                    ))}
                  </SelectContent>
                </Select>
              </FormControl>

              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="format"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("domain.deployment.form.file_format.label")}</FormLabel>
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder={t("domain.deployment.form.file_format.placeholder")} />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem value="pem">PEM</SelectItem>
                  <SelectItem value="pfx">PFX</SelectItem>
                  <SelectItem value="jks">JKS</SelectItem>
                </SelectContent>
              </Select>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="certPath"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("domain.deployment.form.file_cert_path.label")}</FormLabel>
              <FormControl>
                <Input placeholder={t("domain.deployment.form.file_cert_path.label")} {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="keyPath"
          render={({ field }) => (
            <FormItem>
              <FormLabel>密钥路径</FormLabel>
              <FormControl>
                <Input placeholder="输入密钥路径" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        {format === "pfx" && (
          <FormField
            control={form.control}
            name="pfxPassword"
            render={({ field }) => (
              <FormItem>
                <FormLabel>PFX 密码</FormLabel>
                <FormControl>
                  <Input type="password" placeholder="输入 PFX 密码" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        )}

        {format === "jks" && (
          <>
            <FormField
              control={form.control}
              name="jksAlias"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>JKS 别名</FormLabel>
                  <FormControl>
                    <Input placeholder="输入 JKS 别名" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="jksKeypass"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>JKS Keypass</FormLabel>
                  <FormControl>
                    <Input type="password" placeholder="输入 JKS Keypass" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="jksStorepass"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>JKS Storepass</FormLabel>
                  <FormControl>
                    <Input type="password" placeholder="输入 JKS Storepass" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </>
        )}

        <FormField
          control={form.control}
          name="preCommand"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("domain.deployment.form.shell_pre_command.label")}</FormLabel>
              <FormControl>
                <Textarea placeholder={t("domain.deployment.form.shell_pre_command.placeholder")} {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <FormField
          control={form.control}
          name="command"
          render={({ field }) => (
            <FormItem>
              <FormLabel>{t("domain.deployment.form.shell_command.label")}</FormLabel>
              <FormControl>
                <Textarea placeholder={t("domain.deployment.form.shell_command.placeholder")} {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="flex justify-end">
          <Button type="submit">{t("common.button.save")}</Button>
        </div>
      </form>
    </Form>
  );
};

export default DeployToSSH;
