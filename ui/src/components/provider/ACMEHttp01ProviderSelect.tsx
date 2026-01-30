import { useMemo } from "react";
import { useTranslation } from "react-i18next";
import { Avatar, Select, Typography, theme } from "antd";

import { type ACMEHttp01Provider, acmeHttp01ProvidersMap } from "@/domain/provider";

import { type SharedSelectProps, useSelectDataSource } from "./_shared";

export interface ACMEHttp01ProviderSelectProps extends SharedSelectProps<ACMEHttp01Provider> {
  showAvailability?: boolean;
}

const ACMEHttp01ProviderSelect = ({ showAvailability, onFilter, ...props }: ACMEHttp01ProviderSelectProps) => {
  const { t } = useTranslation();

  const { token: themeToken } = theme.useToken();

  const dataSources = useSelectDataSource({
    dataSource: Array.from(acmeHttp01ProvidersMap.values()),
    filters: [onFilter!],
  });
  const options = useMemo(() => {
    const convert = (providers: ACMEHttp01Provider[]): Array<{ key: string; value: string; label: string; data: ACMEHttp01Provider }> => {
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
    const provider = acmeHttp01ProvidersMap.get(key);
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
      filterOption={(inputValue, option) => {
        if (!option) return false;
        if (!option.label) return false;
        if (!option.value) return false;

        const value = inputValue.toLowerCase();
        return String(option.value).toLowerCase().includes(value) || String(option.label).toLowerCase().includes(value);
      }}
      labelRender={({ value }) => {
        if (value != null) {
          return renderOption(value as string);
        }

        return <span style={{ color: themeToken.colorTextPlaceholder }}>{props.placeholder}</span>;
      }}
      options={options}
      optionFilterProp={void 0}
      optionLabelProp={void 0}
      optionRender={(option) => renderOption(option.data.value as string)}
    />
  );
};

export default ACMEHttp01ProviderSelect;
