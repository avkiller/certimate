import { useTranslation } from "react-i18next";
import { IconBook, IconBrandGithub, IconBrandTelegram, IconCoin, IconMessageChatbot } from "@tabler/icons-react";
import { Badge, Button, Divider, List, Tooltip, Typography } from "antd";

import { APP_DOCUMENT_URL, APP_DOWNLOAD_URL, APP_REPO_URL, APP_VERSION } from "@/domain/app";
import { useVersionChecker } from "@/hooks";

const SettingsAbout = () => {
  const { t } = useTranslation();

  const { hasUpdate } = useVersionChecker();

  const handleDownloadClick = () => {
    window.open(APP_DOWNLOAD_URL, "_blank");
  };

  const handleDocumentClick = () => {
    window.open(APP_DOCUMENT_URL, "_blank");
  };

  const handleGithubClick = () => {
    window.open(APP_REPO_URL, "_blank");
  };

  const handleTelegramClick = () => {
    window.open("https://t.me/+ZXphsppxUg41YmVl", "_blank");
  };

  const handleDonateClick = () => {
    window.open("https://profile.ikit.fun/sponsors/", "_blank");
  };

  const handleFeedbackClick = () => {
    window.open(APP_REPO_URL + "/issues", "_blank");
  };

  return (
    <>
      <h2>Certimate</h2>
      <div className="mb-4">
        <div className="flex items-center gap-2">
          <Typography.Text type="secondary">Version: {APP_VERSION}</Typography.Text>
          <Badge className="cursor-pointer" count={hasUpdate ? t("settings.about.version.new") : void 0} onClick={handleDownloadClick} />
        </div>
      </div>
      <div className="mb-2 flex flex-wrap items-center gap-2">
        <Tooltip title={t("settings.about.socials.document")}>
          <Button type="text" icon={<IconBook size="1.5em" onClick={handleDocumentClick} />} />
        </Tooltip>
        <Tooltip title={t("settings.about.socials.github")}>
          <Button type="text" icon={<IconBrandGithub size="1.5em" onClick={handleGithubClick} />} />
        </Tooltip>
        <Tooltip title={t("settings.about.socials.telegram")}>
          <Button type="text" icon={<IconBrandTelegram size="1.5em" onClick={handleTelegramClick} />} />
        </Tooltip>
        <Tooltip title={t("settings.about.socials.donate")}>
          <Button type="text" icon={<IconCoin size="1.5em" onClick={handleDonateClick} />} />
        </Tooltip>
      </div>

      <Divider />

      <h2>{t("settings.about.contributors.title")}</h2>
      <div className="mb-4">
        <Typography.Text type="secondary">{t("settings.about.contributors.tips")}</Typography.Text>
      </div>
      <div className="mb-2 md:max-w-160">
        <img className="max-w-full" src="https://contrib.rocks/image?repo=certimate-go/certimate" alt="Contributors" />
      </div>

      <Divider />

      <div className="md:max-w-160">
        <List bordered>
          <List.Item extra={<Button onClick={handleFeedbackClick}>{t("settings.about.feedback.button")}</Button>}>
            <List.Item.Meta
              avatar={<IconMessageChatbot size="1.5em" />}
              title={t("settings.about.feedback.title")}
              description={t("settings.about.feedback.subtitle")}
            />
          </List.Item>
        </List>
      </div>
    </>
  );
};

export default SettingsAbout;
