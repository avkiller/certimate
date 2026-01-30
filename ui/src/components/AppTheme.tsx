import { useTranslation } from "react-i18next";
import { IconMoon, type IconProps, IconSun, IconSunMoon } from "@tabler/icons-react";
import { Dropdown, type DropdownProps, Typography } from "antd";

import Show from "@/components/Show";
import { useBrowserTheme } from "@/hooks";
import { mergeCls } from "@/utils/css";

export const useAppThemeMenuItems = () => {
  const { t } = useTranslation();

  const { themeMode, setThemeMode } = useBrowserTheme();

  const items = (
    [
      ["light", "common.theme.light", <IconSun size="1em" />],
      ["dark", "common.theme.dark", <IconMoon size="1em" />],
      ["system", "common.theme.system", <IconSunMoon size="1em" />],
    ] satisfies Array<[string, string, React.ReactNode]>
  ).map(([key, label, icon]) => {
    return {
      key: key,
      label: t(label),
      icon: icon,
      onClick: () => {
        if (key !== themeMode) {
          setThemeMode(key as Parameters<typeof setThemeMode>[0]);
          window.location.reload();
        }
      },
    };
  });

  return items;
};

export interface AppThemeDropdownProps {
  children?: React.ReactNode;
  trigger?: DropdownProps["trigger"];
}

const AppThemeDropdown = ({ children, trigger = ["click"] }: AppThemeDropdownProps) => {
  const items = useAppThemeMenuItems();

  return (
    <Dropdown menu={{ items }} trigger={trigger}>
      {children}
    </Dropdown>
  );
};

export interface AppThemeIconProps extends IconProps {}

const AppThemeIcon = (props: AppThemeIconProps) => {
  const { theme } = useBrowserTheme();

  return (
    <Show>
      <Show.Case when={theme === "dark"}>
        <IconMoon {...props} />
      </Show.Case>
      <Show.Default>
        <IconSun {...props} />
      </Show.Default>
    </Show>
  );
};

export interface AppThemeLinkButtonProps {
  className?: string;
  style?: React.CSSProperties;
  showIcon?: boolean;
}

const AppThemeLinkButton = ({ className, style, showIcon = true }: AppThemeLinkButtonProps) => {
  const { t } = useTranslation();

  const { themeMode } = useBrowserTheme();

  return (
    <AppThemeDropdown trigger={["click", "hover"]}>
      <Typography.Text className={mergeCls("cursor-pointer", className)} style={style} type="secondary">
        <div className="flex items-center justify-center space-x-1">
          {showIcon ? <AppThemeIcon size="1em" /> : <></>}
          <span>{t(`common.theme.${themeMode}`)}</span>
        </div>
      </Typography.Text>
    </AppThemeDropdown>
  );
};

export default {
  Dropdown: AppThemeDropdown,
  Icon: AppThemeIcon,
  LinkButton: AppThemeLinkButton,
};
