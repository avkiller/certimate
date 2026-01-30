import { useTranslation } from "react-i18next";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";
import { Form } from "antd";

import { NodeConfigDrawer } from "./_shared";
import BizUploadNodeConfigForm from "./BizUploadNodeConfigForm";
import { NodeType } from "../nodes/typings";

export interface BizUploadNodeConfigDrawerProps {
  afterClose?: () => void;
  loading?: boolean;
  node: FlowNodeEntity;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
}

const BizUploadNodeConfigDrawer = ({ node, ...props }: BizUploadNodeConfigDrawerProps) => {
  if (node.flowNodeType !== NodeType.BizUpload) {
    console.warn(`[certimate] current workflow node type is not: ${NodeType.BizUpload}`);
  }

  const { i18n } = useTranslation();

  const [formInst] = Form.useForm();

  return (
    <NodeConfigDrawer
      anchor={{
        items: BizUploadNodeConfigForm.getAnchorItems({ i18n }),
      }}
      form={formInst}
      node={node}
      {...props}
    >
      <BizUploadNodeConfigForm form={formInst} node={node} />
    </NodeConfigDrawer>
  );
};

export default BizUploadNodeConfigDrawer;
