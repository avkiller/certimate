import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";

export interface NullProps {
  node: FlowNodeEntity;
}

const Null = (_: NullProps) => {
  return null;
};

export default Null;
