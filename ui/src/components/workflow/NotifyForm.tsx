import { WorkflowNode, WorkflowNodeConfig } from "@/domain/workflow";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "../ui/form";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger } from "../ui/select";
import { Input } from "../ui/input";
import { useWorkflowStore, WorkflowState } from "@/stores/workflow";
import { useShallow } from "zustand/shallow";
import { usePanel } from "./PanelProvider";
import { useTranslation } from "react-i18next";
import { Button } from "../ui/button";
import { useNotifyContext } from "@/providers/notify";
import { useEffect, useState } from "react";
import { NotifyChannels, channels as supportedChannels } from "@/domain/settings";
import { SelectValue } from "@radix-ui/react-select";
import { Textarea } from "../ui/textarea";
import { RefreshCw, Settings } from "lucide-react";

type NotifyFormProps = {
  data: WorkflowNode;
};

const selectState = (state: WorkflowState) => ({
  updateNode: state.updateNode,
});
type ChannelName = {
  name: string;
  label: string;
};

const i18nPrefix = "workflow.node.notify.form";
const NotifyForm = ({ data }: NotifyFormProps) => {
  const { updateNode } = useWorkflowStore(useShallow(selectState));
  const { hidePanel } = usePanel();
  const { t } = useTranslation();
  const { config: notifyConfig, initChannels } = useNotifyContext();

  const [chanels, setChanels] = useState<ChannelName[]>([]);

  useEffect(() => {
    setChanels(getChannels());
  }, [notifyConfig]);

  const getChannels = () => {
    const rs: ChannelName[] = [];
    if (!notifyConfig.content) {
      return rs;
    }

    const chanels = notifyConfig.content as NotifyChannels;
    for (const channel of supportedChannels) {
      if (chanels[channel.name] && chanels[channel.name].enabled) {
        rs.push(channel);
      }
    }
    return rs;
  };

  const formSchema = z.object({
    channel: z.string(),
    title: z.string().min(1),
    content: z.string().min(1),
  });

  let config: WorkflowNodeConfig = {
    channel: "",
    title: "",
    content: "",
  };

  if (data) config = data.config ?? config;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      channel: config.channel as string,
      title: config.title as string,
      content: config.content as string,
    },
  });

  const onSubmit = (config: z.infer<typeof formSchema>) => {
    updateNode({ ...data, config, validated: true });
    hidePanel();
  };

  return (
    <>
      <Form {...form}>
        <form
          onSubmit={(e) => {
            e.stopPropagation();
            form.handleSubmit(onSubmit)(e);
          }}
          className="space-y-8 dark:text-stone-200"
        >
          <FormField
            control={form.control}
            name="channel"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="flex justify-between items-center">
                  <div className="flex space-x-2 items-center">
                    <div>{t(`${i18nPrefix}.channel.label`)}</div>
                    <RefreshCw size={16} className="cursor-pointer" onClick={() => initChannels()} />
                  </div>
                  <a
                    href="#/setting/notify"
                    target="_blank"
                    className="flex justify-between items-center space-x-1 font-normal text-primary hover:underline cursor-pointer"
                  >
                    <Settings size={16} /> <div>{t(`${i18nPrefix}.settingChannel.label`)}</div>
                  </a>
                </FormLabel>
                <FormControl>
                  <Select
                    {...field}
                    value={field.value}
                    onValueChange={(value) => {
                      form.setValue("channel", value);
                    }}
                  >
                    <SelectTrigger>
                      <SelectValue placeholder={t(`${i18nPrefix}.channel.placeholder`)} />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        {chanels.map((item) => (
                          <SelectItem key={item.name} value={item.name}>
                            <div>{t(item.label)}</div>
                          </SelectItem>
                        ))}
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t(`${i18nPrefix}.title.label`)}</FormLabel>
                <FormControl>
                  <Input placeholder={t(`${i18nPrefix}.title.placeholder`)} {...field} />
                </FormControl>

                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="content"
            render={({ field }) => (
              <FormItem>
                <FormLabel>{t(`${i18nPrefix}.content.label`)}</FormLabel>
                <FormControl>
                  <Textarea placeholder={t(`${i18nPrefix}.content.placeholder`)} {...field} />
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

export default NotifyForm;
