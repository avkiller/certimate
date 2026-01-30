import { useTranslation } from "react-i18next";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";
import { Form } from "antd";

import { NodeConfigDrawer } from "./_shared";
import DelayNodeConfigForm from "./DelayNodeConfigForm";
import { NodeType } from "../nodes/typings";

export interface DelayNodeConfigDrawerProps {
  afterClose?: () => void;
  loading?: boolean;
  node: FlowNodeEntity;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
}

const DelayNodeConfigDrawer = ({ node, ...props }: DelayNodeConfigDrawerProps) => {
  if (node.flowNodeType !== NodeType.Delay) {
    console.warn(`[certimate] current workflow node type is not: ${NodeType.Delay}`);
  }

  const { i18n } = useTranslation();

  const [formInst] = Form.useForm();

  return (
    <NodeConfigDrawer
      anchor={{
        items: DelayNodeConfigForm.getAnchorItems({ i18n }),
      }}
      form={formInst}
      node={node}
      {...props}
    >
      <DelayNodeConfigForm form={formInst} node={node} />
    </NodeConfigDrawer>
  );
};

export default DelayNodeConfigDrawer;
