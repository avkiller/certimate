import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { useToast } from "@/components/ui/use-toast";
import { getErrMsg } from "@/utils/error";
import { isValidURL } from "@/utils/url";
import { NotifyChannels, NotifyChannelServerChan } from "@/domain/settings";
import { save } from "@/repository/settings";
import { useNotifyContext } from "@/providers/notify";
import { notifyTest } from "@/api/notify";
import Show from "@/components/Show";

type ServerChanSetting = {
  id: string;
  name: string;
  data: NotifyChannelServerChan;
};

const ServerChan = () => {
  const { config, setChannels } = useNotifyContext();
  const { t } = useTranslation();
  const [changed, setChanged] = useState<boolean>(false);

  const [serverchan, setServerChan] = useState<ServerChanSetting>({
    id: config.id ?? "",
    name: "notifyChannels",
    data: {
      url: "",
      enabled: false,
    },
  });

  const [originServerChan, setOriginServerChan] = useState<ServerChanSetting>({
    id: config.id ?? "",
    name: "notifyChannels",
    data: {
      url: "",
      enabled: false,
    },
  });

  useEffect(() => {
    setChanged(false);
  }, [config]);

  useEffect(() => {
    const data = getDetailServerChan();
    setOriginServerChan({
      id: config.id ?? "",
      name: "serverchan",
      data,
    });
  }, [config]);

  useEffect(() => {
    const data = getDetailServerChan();
    setServerChan({
      id: config.id ?? "",
      name: "serverchan",
      data,
    });
  }, [config]);

  const { toast } = useToast();

  const checkChanged = (data: NotifyChannelServerChan) => {
    if (data.url !== originServerChan.data.url) {
      setChanged(true);
    } else {
      setChanged(false);
    }
  };

  const getDetailServerChan = () => {
    const df: NotifyChannelServerChan = {
      url: "",
      enabled: false,
    };
    if (!config.content) {
      return df;
    }
    const chanels = config.content as NotifyChannels;
    if (!chanels.serverchan) {
      return df;
    }

    return chanels.serverchan as NotifyChannelServerChan;
  };

  const handleSaveClick = async () => {
    try {
      serverchan.data.url = serverchan.data.url.trim();
      if (!isValidURL(serverchan.data.url)) {
        toast({
          title: t("common.text.operation_failed"),
          description: t("common.errmsg.url_invalid"),
          variant: "destructive",
        });
        return;
      }

      const resp = await save({
        ...config,
        name: "notifyChannels",
        content: {
          ...config.content,
          serverchan: {
            ...serverchan.data,
          },
        },
      });

      setChannels(resp);
      toast({
        title: t("common.text.operation_succeeded"),
        description: t("settings.notification.config.saved.message"),
      });
    } catch (e) {
      const msg = getErrMsg(e);

      toast({
        title: t("common.text.operation_failed"),
        description: `${t("settings.notification.config.failed.message")}: ${msg}`,
        variant: "destructive",
      });
    }
  };

  const [testing, setTesting] = useState<boolean>(false);
  const handlePushTestClick = async () => {
    if (testing) return;

    try {
      setTesting(true);

      await notifyTest("serverchan");

      toast({
        title: t("settings.notification.push_test_message.succeeded.message"),
        description: t("settings.notification.push_test_message.succeeded.message"),
      });
    } catch (e) {
      const msg = getErrMsg(e);

      toast({
        title: t("settings.notification.push_test_message.failed.message"),
        description: `${t("settings.notification.push_test_message.failed.message")}: ${msg}`,
        variant: "destructive",
      });
    } finally {
      setTesting(false);
    }
  };

  const handleSwitchChange = async () => {
    const newData = {
      ...serverchan,
      data: {
        ...serverchan.data,
        enabled: !serverchan.data.enabled,
      },
    };
    setServerChan(newData);

    try {
      const resp = await save({
        ...config,
        name: "notifyChannels",
        content: {
          ...config.content,
          serverchan: {
            ...newData.data,
          },
        },
      });

      setChannels(resp);
    } catch (e) {
      const msg = getErrMsg(e);

      toast({
        title: t("common.text.operation_failed"),
        description: `${t("settings.notification.config.failed.message")}: ${msg}`,
        variant: "destructive",
      });
    }
  };

  return (
    <div className="flex flex-col space-y-4">
      <div>
        <Label>{t("settings.notification.serverchan.url.label")}</Label>
        <Input
          placeholder={t("settings.notification.serverchan.url.placeholder")}
          value={serverchan.data.url}
          onChange={(e) => {
            const newData = {
              ...serverchan,
              data: {
                ...serverchan.data,
                url: e.target.value,
              },
            };

            checkChanged(newData.data);
            setServerChan(newData);
          }}
        />
      </div>

      <div className="flex justify-between gap-4">
        <div className="flex items-center space-x-1">
          <Switch id="airplane-mode" checked={serverchan.data.enabled} onCheckedChange={handleSwitchChange} />
          <Label htmlFor="airplane-mode">{t("settings.notification.config.enable")}</Label>
        </div>

        <div className="flex items-center space-x-1">
          <Show when={changed}>
            <Button
              onClick={() => {
                handleSaveClick();
              }}
            >
              {t("common.button.save")}
            </Button>
          </Show>

          <Show when={!changed && serverchan.id != ""}>
            <Button
              variant="secondary"
              loading={testing}
              onClick={() => {
                handlePushTestClick();
              }}
            >
              {t("settings.notification.push_test_message")}
            </Button>
          </Show>
        </div>
      </div>
    </div>
  );
};

export default ServerChan;
