import { useEffect } from "react";
import { type NodeRenderProps, useClientContext, useNodeRender, useRefresh } from "@flowgram.ai/fixed-layout-editor";

import { useDesignerContext } from "./_context";
import { NodeRenderContextProvider } from "./NodeRenderContext";
import { type NodeRegistry } from "./nodes/typings";

export interface NodeProps extends NodeRenderProps {}

const Node = (_: NodeProps) => {
  const ctx = useClientContext();

  const refresh = useRefresh();

  const nodeRender = useNodeRender();

  const { onDocumentChange: fireOnDocumentChange, onNodeChange: fireOnNodeChange, onNodeClick: fireOnNodeClick } = useDesignerContext();

  useEffect(() => {
    const d = ctx.document.originTree.onTreeChange(() => refresh());

    return () => d.dispose();
  }, []);

  useEffect(() => {
    const d1 = nodeRender.form?.onFormValuesChange?.(() => {
      refresh();

      fireOnNodeChange(nodeRender.node);
      fireOnDocumentChange();
    });
    const d2 = nodeRender.form?.onValidate?.(() => {
      refresh();
    });

    return () => {
      d1?.dispose();
      d2?.dispose();
    };
  }, [nodeRender.form]);

  const handleNodeClick = () => {
    const node = nodeRender.node;
    if (node.getNodeRegistry<NodeRegistry>().meta?.clickable) {
      fireOnNodeClick(node);
    }
  };

  return (
    <div
      style={{
        opacity: nodeRender.dragging ? 0.3 : 1,
        ...nodeRender.node.getNodeRegistry<NodeRegistry>().meta?.style,
      }}
      onMouseEnter={nodeRender.onMouseEnter}
      onMouseLeave={nodeRender.onMouseLeave}
      onMouseDown={(e) => {
        nodeRender.startDrag(e);
        e.stopPropagation();
      }}
      onClick={handleNodeClick}
    >
      <NodeRenderContextProvider value={nodeRender}>{nodeRender.form?.render()}</NodeRenderContextProvider>
    </div>
  );
};

export default Node;
