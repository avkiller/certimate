import { useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Avatar, Select, Typography, theme } from "antd";

import { type NotificationProvider, notificationProvidersMap } from "@/domain/provider";
import { matchSearchOption } from "@/utils/search";

import { type SharedSelectProps, useSelectDataSource } from "./_shared";

export interface NotificationProviderSelectProps extends SharedSelectProps<NotificationProvider> {
  showAvailability?: boolean;
}

const NotificationProviderSelect = ({ showAvailability = false, onFilter, ...props }: NotificationProviderSelectProps) => {
  const { t } = useTranslation();

  const { token: themeToken } = theme.useToken();

  const dataSources = useSelectDataSource({
    dataSource: Array.from(notificationProvidersMap.values()),
    filters: [onFilter!],
  });
  const options = useMemo(() => {
    const convert = (providers: NotificationProvider[]): Array<{ key: string; value: string; label: string; data: NotificationProvider }> => {
      return providers.map((provider) => ({
        key: provider.type,
        value: provider.type,
        label: t(provider.name),
        data: provider,
      }));
    };

    return showAvailability
      ? [
          {
            label: t("provider.text.available_group"),
            options: convert(dataSources.available),
          },
          {
            label: t("provider.text.unavailable_group"),
            options: convert(dataSources.unavailable),
          },
        ].filter((group) => group.options.length > 0)
      : convert(dataSources.filtered);
  }, [showAvailability, dataSources]);

  const renderOption = (key: string) => {
    const provider = notificationProvidersMap.get(key);
    return (
      <div className="flex items-center gap-2 truncate overflow-hidden">
        <Avatar shape="square" src={provider?.icon} size="small" />
        <Typography.Text ellipsis>{t(provider?.name ?? "")}</Typography.Text>
      </div>
    );
  };

  return (
    <Select
      {...props}
      labelRender={({ value }) => {
        if (value != null) {
          return renderOption(value as string);
        }

        return <span style={{ color: themeToken.colorTextPlaceholder }}>{props.placeholder}</span>;
      }}
      options={options}
      optionLabelProp={void 0}
      optionRender={(option) => renderOption(option.data.value as string)}
      showSearch={{
        filterOption: (inputValue, option) => matchSearchOption(inputValue, option!),
      }}
    />
  );
};

export default NotificationProviderSelect;
