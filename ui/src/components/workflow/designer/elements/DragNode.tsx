import { type FlowNodeEntity, type DragNodeProps as FlowgramDragNodeProps } from "@flowgram.ai/fixed-layout-editor";
import { Badge, Card } from "antd";

export interface DragNodeProps extends FlowgramDragNodeProps {
  dragStart: FlowNodeEntity;
  dragNodes: FlowNodeEntity[];
}

const DragNode = ({ dragStart, dragNodes }: DragNodeProps) => {
  const count = (dragNodes || [])
    .map((n) => (n.allCollapsedChildren.length ? n.allCollapsedChildren.filter((_n) => !_n.hidden).length : 1))
    .reduce((acc, cur) => acc + cur, 0);
  return (
    <Badge count={count > 1 ? count : 0} size="small">
      <div className="relative w-[160px]">
        <Card className="bg-transparent shadow" styles={{ body: { padding: 0 } }}>
          <div className="overflow-hidden px-4 py-2 text-primary">
            <div className="truncate">{dragStart ? dragStart.form?.getValueIn("name") || `#${dragStart?.id}` : "\u00A0"}</div>
          </div>
        </Card>
      </div>
    </Badge>
  );
};

export default DragNode;
