import { useTranslation } from "react-i18next";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";
import { Form } from "antd";

import { NodeConfigDrawer } from "./_shared";
import StartNodeConfigForm from "./StartNodeConfigForm";
import { NodeType } from "../nodes/typings";

export interface StartNodeConfigDrawerProps {
  afterClose?: () => void;
  loading?: boolean;
  node: FlowNodeEntity;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
}

const StartNodeConfigDrawer = ({ node, ...props }: StartNodeConfigDrawerProps) => {
  if (node.flowNodeType !== NodeType.Start) {
    console.warn(`[certimate] current workflow node type is not: ${NodeType.Start}`);
  }

  const { i18n } = useTranslation();

  const [formInst] = Form.useForm();

  return (
    <NodeConfigDrawer
      anchor={{
        items: StartNodeConfigForm.getAnchorItems({ i18n }),
      }}
      form={formInst}
      node={node}
      {...props}
    >
      <StartNodeConfigForm form={formInst} node={node} />
    </NodeConfigDrawer>
  );
};

export default StartNodeConfigDrawer;
