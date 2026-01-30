import { useTranslation } from "react-i18next";
import { IconBook } from "@tabler/icons-react";
import { Typography } from "antd";

import { APP_DOCUMENT_URL } from "@/domain/app";

export interface AppDocumentLinkButtonProps {
  className?: string;
  style?: React.CSSProperties;
  showIcon?: boolean;
}

const AppDocumentLinkButton = ({ className, style, showIcon = true }: AppDocumentLinkButtonProps) => {
  const { t } = useTranslation();

  const handleDocumentClick = () => {
    window.open(APP_DOCUMENT_URL, "_blank");
  };

  return (
    <Typography.Link className={className} style={style} type="secondary" onClick={handleDocumentClick}>
      <div className="flex items-center justify-center space-x-1">
        {showIcon ? <IconBook size="1em" /> : <></>}
        <span>{t("common.menu.document")}</span>
      </div>
    </Typography.Link>
  );
};

export default {
  LinkButton: AppDocumentLinkButton,
};
