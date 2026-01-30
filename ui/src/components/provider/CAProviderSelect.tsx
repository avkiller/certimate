import { useMemo } from "react";
import { useTranslation } from "react-i18next";
import { useControllableValue } from "ahooks";
import { Avatar, Select, Typography, theme } from "antd";

import { type CAProvider, caProvidersMap } from "@/domain/provider";

import { type SharedSelectProps, useSelectDataSource } from "./_shared";

export interface CAProviderSelectProps extends SharedSelectProps<CAProvider> {
  showAvailability?: boolean;
  showDefault?: boolean;
}

const CAProviderSelect = ({ showAvailability, showDefault, onFilter, ...props }: CAProviderSelectProps) => {
  const { t } = useTranslation();

  const { token: themeToken } = theme.useToken();

  const [value, setValue] = useControllableValue<string | undefined>(props, {
    valuePropName: "value",
    defaultValuePropName: "defaultValue",
    trigger: "onChange",
  });

  const dataSources = useSelectDataSource({
    dataSource: Array.from(caProvidersMap.values()),
    filters: [onFilter!],
  });
  const options = useMemo(() => {
    const convert = (providers: CAProvider[]): Array<{ key: string; value: string; label: string; data: CAProvider }> => {
      return providers.map((provider) => ({
        key: provider.type,
        value: provider.type,
        label: t(provider.name),
        data: provider,
      }));
    };

    const defaultOption = {
      key: "",
      value: "",
      data: {} as CAProvider,
    };
    const plainOptions = convert(dataSources.filtered);
    const groupOptions = [
      {
        label: t("provider.text.available_group"),
        options: convert(dataSources.available),
      },
      {
        label: t("provider.text.unavailable_group"),
        options: convert(dataSources.unavailable),
      },
    ].filter((group) => group.options.length > 0);

    const temp = showAvailability
      ? showDefault
        ? [{ label: t("provider.text.default_group"), options: [defaultOption] }, ...groupOptions]
        : groupOptions
      : showDefault
        ? [defaultOption, ...plainOptions]
        : plainOptions;

    return temp;
  }, [showAvailability, showDefault, dataSources]);

  const renderOption = (key: string) => {
    if (key === "") {
      return (
        <div className="truncate">
          <Typography.Text ellipsis>{showAvailability ? t("provider.text.default_ca_in_group") : t("provider.text.default_ca")}</Typography.Text>
        </div>
      );
    }

    const provider = caProvidersMap.get(key);
    return (
      <div className="flex items-center gap-2 truncate overflow-hidden">
        <Avatar shape="square" src={provider?.icon} size="small" />
        <Typography.Text ellipsis>{t(provider?.name ?? "")}</Typography.Text>
      </div>
    );
  };

  const handleChange = (value: string) => {
    setValue((_) => (value !== "" ? value : void 0));
  };

  return (
    <Select
      {...props}
      filterOption={(inputValue, option) => {
        if (option?.value === "") return true;
        if (!option) return false;
        if (!option.label) return false;
        if (!option.value) return false;

        const value = inputValue.toLowerCase();
        return String(option.value).toLowerCase().includes(value) || String(option.label).toLowerCase().includes(value);
      }}
      labelRender={({ value }) => {
        if (value != null && value !== "") {
          return renderOption(value as string);
        }

        return <span style={{ color: themeToken.colorTextPlaceholder }}>{props.placeholder}</span>;
      }}
      options={options}
      optionFilterProp={void 0}
      optionLabelProp={void 0}
      optionRender={(option) => renderOption(option.data.value as string)}
      value={value}
      onChange={handleChange}
      onSelect={handleChange}
    />
  );
};

export default CAProviderSelect;
