import { useTranslation } from "react-i18next";
import { IconLanguage, type IconProps } from "@tabler/icons-react";
import { Dropdown, type DropdownProps, Typography } from "antd";

import { IconLanguageEnZh, IconLanguageZhEn } from "@/components/icons";
import Show from "@/components/Show";
import { localeNames, localeResources } from "@/i18n";
import { mergeCls } from "@/utils/css";

export const useAppLocaleMenuItems = () => {
  const { i18n } = useTranslation();

  const items = Object.keys(i18n.store.data).map((key) => {
    return {
      key: key as string,
      label: i18n.store.data[key].name as string,
      onClick: () => {
        if (key !== (i18n.resolvedLanguage ?? i18n.language)) {
          i18n.changeLanguage(key);
          window.location.reload();
        }
      },
    };
  });

  return items;
};

export interface AppLocaleDropdownProps {
  children?: React.ReactNode;
  trigger?: DropdownProps["trigger"];
}

const AppLocaleDropdown = ({ children, trigger = ["click"] }: AppLocaleDropdownProps) => {
  const items = useAppLocaleMenuItems();

  return (
    <Dropdown menu={{ items }} trigger={trigger}>
      {children}
    </Dropdown>
  );
};

export interface AppLocaleIconProps extends IconProps {}

const AppLocaleIcon = (props: AppLocaleIconProps) => {
  const { i18n } = useTranslation();

  return (
    <Show>
      <Show.Case when={(i18n.resolvedLanguage ?? i18n.language) === localeNames.EN}>
        <IconLanguageEnZh {...props} />
      </Show.Case>
      <Show.Case when={(i18n.resolvedLanguage ?? i18n.language) === localeNames.ZH}>
        <IconLanguageZhEn {...props} />
      </Show.Case>
      <Show.Default>
        <IconLanguage {...props} />
      </Show.Default>
    </Show>
  );
};

export interface AppLocaleLinkButtonProps {
  className?: string;
  style?: React.CSSProperties;
  showIcon?: boolean;
}

const AppLocaleLinkButton = ({ className, style, showIcon = true }: AppLocaleLinkButtonProps) => {
  const { t } = useTranslation();
  const { i18n } = useTranslation();

  return (
    <AppLocaleDropdown trigger={["click", "hover"]}>
      <Typography.Text className={mergeCls("cursor-pointer", className)} style={style} type="secondary">
        <div className="flex items-center justify-center space-x-1">
          {showIcon ? <AppLocaleIcon size="1em" /> : <></>}
          <span>{String(localeResources[i18n.resolvedLanguage ?? i18n.language]?.name ?? t("common.menu.locale"))}</span>
        </div>
      </Typography.Text>
    </AppLocaleDropdown>
  );
};

export default {
  Dropdown: AppLocaleDropdown,
  Icon: AppLocaleIcon,
  LinkButton: AppLocaleLinkButton,
};
