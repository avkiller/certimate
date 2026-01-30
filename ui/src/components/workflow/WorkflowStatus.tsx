import { useTranslation } from "react-i18next";
import {
  IconCircleCheck,
  IconCircleCheckFilled,
  IconCircleDashed,
  IconCircleOff,
  IconCircleX,
  IconCircleXFilled,
  IconClock,
  IconClockFilled,
  IconLoader3,
} from "@tabler/icons-react";
import { Typography, theme } from "antd";

import { WORKFLOW_RUN_STATUSES, type WorkflorRunStatusType } from "@/domain/workflowRun";
import { mergeCls } from "@/utils/css";

const useColor = (value: WorkflorRunStatusType | string, defaultColor?: string | false) => {
  const { token: themeToken } = theme.useToken();

  switch (value) {
    case WORKFLOW_RUN_STATUSES.PENDING:
      if (defaultColor == null || !defaultColor) {
        return themeToken.colorTextSecondary;
      }
      break;
    case WORKFLOW_RUN_STATUSES.PROCESSING:
      if (defaultColor == null || !defaultColor) {
        return themeToken.colorInfo;
      }
      break;
    case WORKFLOW_RUN_STATUSES.SUCCEEDED:
      if (defaultColor == null || !defaultColor) {
        return themeToken.colorSuccess;
      }
      break;
    case WORKFLOW_RUN_STATUSES.FAILED:
      if (defaultColor == null || !defaultColor) {
        return themeToken.colorError;
      }
      break;
    case WORKFLOW_RUN_STATUSES.CANCELED:
      if (defaultColor == null || !defaultColor) {
        return themeToken.colorWarning;
      }
      break;
    default:
      if (defaultColor == null || !defaultColor) {
        return themeToken.colorTextSecondary;
      }
      break;
  }

  return defaultColor;
};

export interface WorkflowStatusIconProps {
  className?: string;
  style?: React.CSSProperties;
  color?: string | false;
  size?: number | string;
  type?: "filled" | "outlined";
  value: WorkflorRunStatusType | string;
}

const WorkflowStatusIcon = ({ className, style, size = "1.25em", type = "outlined", value, ...props }: WorkflowStatusIconProps) => {
  const color = useColor(value, props.color);

  switch (value) {
    case WORKFLOW_RUN_STATUSES.PENDING:
      return (
        <span className={mergeCls("anticon", className)} style={style} role="img">
          {type === "filled" ? <IconClockFilled color={color} size={size} /> : <IconClock color={color} size={size} />}
        </span>
      );
    case WORKFLOW_RUN_STATUSES.PROCESSING:
      return (
        <span className={mergeCls("anticon", "animate-spin", className)} style={style} role="img">
          <IconLoader3 color={color} size={size} />
        </span>
      );
    case WORKFLOW_RUN_STATUSES.SUCCEEDED:
      return (
        <span className={mergeCls("anticon", className)} style={style} role="img">
          {type === "filled" ? <IconCircleCheckFilled color={color} size={size} /> : <IconCircleCheck color={color} size={size} />}
        </span>
      );
    case WORKFLOW_RUN_STATUSES.FAILED:
      return (
        <span className={mergeCls("anticon", className)} style={style} role="img">
          {type === "filled" ? <IconCircleXFilled color={color} size={size} /> : <IconCircleX color={color} size={size} />}
        </span>
      );
    case WORKFLOW_RUN_STATUSES.CANCELED:
      return (
        <span className={mergeCls("anticon", className)} style={style} role="img">
          <IconCircleOff color={color} size={size} />
        </span>
      );
    default:
      return (
        <span className={mergeCls("anticon", className)} style={style} role="img">
          <IconCircleDashed color={color} size={size} />
        </span>
      );
  }
};

export interface WorkflowStatusProps {
  className?: string;
  style?: React.CSSProperties;
  children?: React.ReactNode;
  color?: string | false;
  showIcon?: boolean;
  type?: WorkflowStatusIconProps["type"];
  value: WorkflorRunStatusType | string;
}

const WorkflowStatus = ({ className, style, children, showIcon = true, type, value, ...props }: WorkflowStatusProps) => {
  const { t } = useTranslation();

  const color = useColor(value, props.color);

  const renderIcon = () => (showIcon ? <WorkflowStatusIcon type={type} value={value} /> : null);

  switch (value) {
    case WORKFLOW_RUN_STATUSES.PENDING:
    case WORKFLOW_RUN_STATUSES.PROCESSING:
    case WORKFLOW_RUN_STATUSES.SUCCEEDED:
    case WORKFLOW_RUN_STATUSES.FAILED:
    case WORKFLOW_RUN_STATUSES.CANCELED:
      return (
        <Typography.Text className={className} style={style}>
          <div className="flex items-center gap-2">
            {renderIcon()}
            {children != null ? children : <span style={{ color: color }}>{t(`workflow_run.props.status.${value.toLowerCase()}`)}</span>}
          </div>
        </Typography.Text>
      );
    default:
      return (
        <Typography.Text className={className} style={style}>
          <div className="flex items-center gap-2">{children != null ? children : <></>}</div>
        </Typography.Text>
      );
  }
};

const _default = Object.assign(WorkflowStatus, {
  Icon: WorkflowStatusIcon,
});

export default _default;
