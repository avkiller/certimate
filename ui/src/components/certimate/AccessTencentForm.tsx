import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { ClientResponseError } from "pocketbase";

import { Button } from "@/components/ui/button";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { PbErrorData } from "@/domain/base";
import { accessProvidersMap, accessTypeFormSchema, type AccessModel, type TencentConfig } from "@/domain/access";
import { save } from "@/repository/access";
import { useAccessStore } from "@/stores/access";

type AccessTencentFormProps = {
  op: "add" | "edit" | "copy";
  data?: AccessModel;
  onAfterReq: () => void;
};

const AccessTencentForm = ({ data, op, onAfterReq }: AccessTencentFormProps) => {
  const { createAccess, updateAccess } = useAccessStore();
  const { t } = useTranslation();
  const formSchema = z.object({
    id: z.string().optional(),
    name: z
      .string()
      .min(1, "access.authorization.form.name.placeholder")
      .max(64, t("common.errmsg.string_max", { max: 64 })),
    configType: accessTypeFormSchema,
    secretId: z
      .string()
      .min(1, "access.authorization.form.secret_id.placeholder")
      .max(64, t("common.errmsg.string_max", { max: 64 })),
    secretKey: z
      .string()
      .min(1, "access.authorization.form.secret_key.placeholder")
      .max(64, t("common.errmsg.string_max", { max: 64 })),
  });

  let config: TencentConfig = {
    secretId: "",
    secretKey: "",
  };
  if (data) config = data.config as TencentConfig;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      id: data?.id,
      name: data?.name || "",
      configType: "tencent",
      secretId: config.secretId,
      secretKey: config.secretKey,
    },
  });

  const onSubmit = async (data: z.infer<typeof formSchema>) => {
    const req: AccessModel = {
      id: data.id as string,
      name: data.name,
      configType: data.configType,
      usage: accessProvidersMap.get(data.configType)!.usage,
      config: {
        secretId: data.secretId,
        secretKey: data.secretKey,
      },
    };

    try {
      req.id = op == "copy" ? "" : req.id;
      const rs = await save(req);

      onAfterReq();

      req.id = rs.id;
      req.created = rs.created;
      req.updated = rs.updated;
      if (data.id && op == "edit") {
        updateAccess(req);
        return;
      }
      createAccess(req);
    } catch (e) {
      const err = e as ClientResponseError;

      Object.entries(err.response.data as PbErrorData).forEach(([key, value]) => {
        form.setError(key as keyof z.infer<typeof formSchema>, {
          type: "manual",
          message: value.message,
        });
      });
    }
  };

  return (
    <>
      <Form {...form}>
        <form
          onSubmit={(e) => {
            e.stopPropagation();
            form.handleSubmit(onSubmit)(e);
          }}
          className="space-y-8"
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t("access.authorization.form.name.label")}</FormLabel>
                <FormControl>
                  <Input placeholder={t("access.authorization.form.name.placeholder")} {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="id"
            render={({ field }) => (
              <FormItem className="hidden">
                <FormLabel>{t("access.authorization.form.config.label")}</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="configType"
            render={({ field }) => (
              <FormItem className="hidden">
                <FormLabel>{t("access.authorization.form.config.label")}</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="secretId"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t("access.authorization.form.secret_id.label")}</FormLabel>
                <FormControl>
                  <Input placeholder={t("access.authorization.form.secret_id.placeholder")} {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="secretKey"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t("access.authorization.form.secret_key.label")}</FormLabel>
                <FormControl>
                  <Input placeholder={t("access.authorization.form.secret_key.placeholder")} {...field} />
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
    </>
  );
};

export default AccessTencentForm;
