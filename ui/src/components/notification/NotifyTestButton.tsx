import { useTranslation } from "react-i18next";
import { useRequest } from "ahooks";
import { Button, type ButtonProps, message, notification } from "antd";

import { notifyTest } from "@/api/notify";
import { getErrMsg } from "@/utils/error";

export interface NotifyTestButtonProps {
  className?: string;
  style?: React.CSSProperties;
  disabled?: boolean;
  provider: string;
  size?: ButtonProps["size"];
}

const NotifyTestButton = ({ className, style, provider, disabled, size }: NotifyTestButtonProps) => {
  const { t } = useTranslation();

  const [messageApi, MessageContextHolder] = message.useMessage();
  const [notificationApi, NotificationContextHolder] = notification.useNotification();

  const { loading, run: executeNotifyTest } = useRequest(
    () => {
      return notifyTest(provider);
    },
    {
      refreshDeps: [provider],
      manual: true,
      onSuccess: () => {
        messageApi.success(t("settings.notification.push_test.pushed"));
      },
      onError: (err) => {
        notificationApi.error({ message: t("common.text.request_error"), description: getErrMsg(err) });
      },
    }
  );

  const handleClick = () => {
    executeNotifyTest();
  };

  return (
    <>
      {MessageContextHolder}
      {NotificationContextHolder}

      <Button className={className} style={style} disabled={disabled} loading={loading} size={size} onClick={handleClick}>
        {t("settings.notification.push_test.button")}
      </Button>
    </>
  );
};

export default NotifyTestButton;
