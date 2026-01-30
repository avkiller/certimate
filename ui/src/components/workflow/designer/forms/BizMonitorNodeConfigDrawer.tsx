import { useTranslation } from "react-i18next";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";
import { Form } from "antd";

import { NodeConfigDrawer } from "./_shared";
import BizMonitorNodeConfigForm from "./BizMonitorNodeConfigForm";
import { NodeType } from "../nodes/typings";

export interface BizMonitorNodeConfigDrawerProps {
  afterClose?: () => void;
  loading?: boolean;
  node: FlowNodeEntity;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
}

const BizMonitorNodeConfigDrawer = ({ node, ...props }: BizMonitorNodeConfigDrawerProps) => {
  if (node.flowNodeType !== NodeType.BizMonitor) {
    console.warn(`[certimate] current workflow node type is not: ${NodeType.BizMonitor}`);
  }

  const { i18n } = useTranslation();

  const [formInst] = Form.useForm();

  return (
    <NodeConfigDrawer
      anchor={{
        items: BizMonitorNodeConfigForm.getAnchorItems({ i18n }),
      }}
      form={formInst}
      node={node}
      {...props}
    >
      <BizMonitorNodeConfigForm form={formInst} node={node} />
    </NodeConfigDrawer>
  );
};

export default BizMonitorNodeConfigDrawer;
