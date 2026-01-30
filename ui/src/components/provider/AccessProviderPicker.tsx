import { useMemo, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { useMount } from "ahooks";
import { Avatar, Card, Empty, Input, type InputRef, Tag, Tooltip, Typography } from "antd";

import Show from "@/components/Show";
import { ACCESS_USAGES, type AccessProvider, type AccessUsageType, accessProvidersMap } from "@/domain/provider";
import { mergeCls } from "@/utils/css";

import { type SharedPickerProps, usePickerDataSource, usePickerWrapperCols } from "./_shared";

export interface AccessProviderPickerProps extends SharedPickerProps<AccessProvider> {
  showOptionTags?: boolean | { [key in AccessUsageType]?: boolean };
}

const AccessProviderPicker = ({
  className,
  style,
  autoFocus,
  gap = "middle",
  placeholder,
  showOptionTags,
  showSearch = false,
  onFilter,
  onSelect,
}: AccessProviderPickerProps) => {
  const { t } = useTranslation();

  const showOptionTagForDNS = useMemo(() => {
    return typeof showOptionTags === "object" ? !!showOptionTags?.[ACCESS_USAGES.DNS] : !!showOptionTags;
  }, [showOptionTags]);
  const showOptionTagForHosting = useMemo(() => {
    return typeof showOptionTags === "object" ? !!showOptionTags?.[ACCESS_USAGES.HOSTING] : !!showOptionTags;
  }, [showOptionTags]);
  const showOptionTagForCA = useMemo(() => {
    return typeof showOptionTags === "object" ? !!showOptionTags?.[ACCESS_USAGES.CA] : !!showOptionTags;
  }, [showOptionTags]);
  const showOptionTagForNotification = useMemo(() => {
    return typeof showOptionTags === "object" ? !!showOptionTags?.[ACCESS_USAGES.NOTIFICATION] : !!showOptionTags;
  }, [showOptionTags]);
  const showOptionTagAnyhow = useMemo(() => {
    return showOptionTagForDNS || showOptionTagForHosting || showOptionTagForCA || showOptionTagForNotification;
  }, [showOptionTagForDNS, showOptionTagForHosting, showOptionTagForCA, showOptionTagForNotification]);

  const { wrapperElRef, cols } = usePickerWrapperCols(showOptionTagAnyhow ? 240 : 200);

  const [keyword, setKeyword] = useState<string>();
  const keywordInputRef = useRef<InputRef>(null);
  useMount(() => {
    if (autoFocus) {
      setTimeout(() => keywordInputRef.current?.focus(), 1);
    }
  });

  const dataSources = usePickerDataSource({
    dataSource: Array.from(accessProvidersMap.values()),
    filters: [onFilter!],
    keyword: keyword,
  });

  const renderOption = (provider: AccessProvider) => {
    return (
      <div key={provider.type}>
        <Card
          className={mergeCls("w-full overflow-hidden shadow", provider.builtin ? " cursor-not-allowed" : "", showOptionTagAnyhow ? "h-32" : "h-28")}
          styles={{ body: { height: "100%", padding: "0.5rem 1rem" } }}
          hoverable
          onClick={() => {
            if (provider.builtin) {
              return;
            }

            handleProviderTypeSelect(provider.type);
          }}
        >
          <div className="flex size-full flex-col items-center justify-center gap-3 overflow-hidden p-2">
            <div className="flex items-center justify-center">
              <Avatar className="bg-stone-100" icon={<img src={provider.icon} />} shape="square" size={32} />
            </div>
            <div className="w-full overflow-hidden text-center">
              <div className={mergeCls("w-full truncate", { "mb-1": showOptionTagAnyhow })}>
                <Tooltip title={t(provider.name)} mouseEnterDelay={1}>
                  <Typography.Text type={provider.builtin ? "secondary" : void 0}>{t(provider.name) || "\u00A0"}</Typography.Text>
                </Tooltip>
              </div>
              <Show when={showOptionTagAnyhow}>
                <div className="origin-top scale-80 whitespace-nowrap" style={{ marginInlineEnd: "-8px" }}>
                  <Show when={provider.builtin}>
                    <Tag>{t("access.props.provider.builtin")}</Tag>
                  </Show>
                  <Show when={showOptionTagForDNS && provider.usages.includes(ACCESS_USAGES.DNS)}>
                    <Tag color="#d93f0b99">{t("access.props.provider.usage.dns")}</Tag>
                  </Show>
                  <Show when={showOptionTagForHosting && provider.usages.includes(ACCESS_USAGES.HOSTING)}>
                    <Tag color="#0052cc99">{t("access.props.provider.usage.hosting")}</Tag>
                  </Show>
                  <Show when={showOptionTagForCA && provider.usages.includes(ACCESS_USAGES.CA)}>
                    <Tag color="#0e8a1699">{t("access.props.provider.usage.ca")}</Tag>
                  </Show>
                  <Show when={showOptionTagForNotification && provider.usages.includes(ACCESS_USAGES.NOTIFICATION)}>
                    <Tag color="#1d76db99">{t("access.props.provider.usage.notification")}</Tag>
                  </Show>
                </div>
              </Show>
            </div>
          </div>
        </Card>
      </div>
    );
  };

  const handleProviderTypeSelect = (value: string) => {
    onSelect?.(value);
  };

  return (
    <div className={className} style={style} ref={wrapperElRef}>
      <Show when={showSearch}>
        <div className="mb-4">
          <Input.Search ref={keywordInputRef} placeholder={placeholder ?? t("common.text.search")} onChange={(e) => setKeyword(e.target.value.trim())} />
        </div>
      </Show>

      <Show when={dataSources.filtered.length > 0} fallback={<Empty description={t("provider.text.nodata")} image={Empty.PRESENTED_IMAGE_SIMPLE} />}>
        <div
          className={mergeCls("grid w-full gap-2", `grid-cols-${cols}`, {
            "gap-4": gap === "large",
            "gap-2": gap === "middle",
            "gap-1": gap === "small",
            [`gap-${+gap || "2"}`]: typeof gap === "number",
          })}
        >
          {dataSources.filtered.map((provider) => renderOption(provider))}
        </div>
      </Show>
    </div>
  );
};

export default AccessProviderPicker;
