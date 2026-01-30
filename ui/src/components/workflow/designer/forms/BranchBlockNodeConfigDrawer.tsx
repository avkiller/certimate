import { useTranslation } from "react-i18next";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";
import { Form } from "antd";

import { NodeConfigDrawer } from "./_shared";
import BranchBlockNodeConfigForm from "./BranchBlockNodeConfigForm";
import { NodeType } from "../nodes/typings";

export interface BranchBlockNodeConfigDrawerProps {
  afterClose?: () => void;
  loading?: boolean;
  node: FlowNodeEntity;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
}

const BranchBlockNodeConfigDrawer = ({ node, ...props }: BranchBlockNodeConfigDrawerProps) => {
  if (node.flowNodeType !== NodeType.BranchBlock) {
    console.warn(`[certimate] current workflow node type is not: ${NodeType.BranchBlock}`);
  }

  const { i18n } = useTranslation();

  const [formInst] = Form.useForm();

  return (
    <NodeConfigDrawer
      anchor={{
        items: BranchBlockNodeConfigForm.getAnchorItems({ i18n }),
      }}
      form={formInst}
      node={node}
      {...props}
    >
      <BranchBlockNodeConfigForm form={formInst} node={node} />
    </NodeConfigDrawer>
  );
};

export default BranchBlockNodeConfigDrawer;
