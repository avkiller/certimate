import { FlowRendererKey } from "@flowgram.ai/fixed-layout-editor";

import Adder from "./Adder";
import BranchAdder from "./BranchAdder";
import Collapse from "./Collapse";
import DraggingAdder from "./DraggingAdder";
import DragHighlightAdder from "./DragHighlightAdder";
import DragNode from "./DragNode";
import Null from "./Null";
import TryCatchCollapse from "./TryCatchCollapse";

export const getAllElements = () => {
  return {
    [FlowRendererKey.ADDER]: Adder,
    [FlowRendererKey.BRANCH_ADDER]: BranchAdder,
    [FlowRendererKey.SLOT_ADDER]: Null,

    [FlowRendererKey.COLLAPSE]: Collapse,
    [FlowRendererKey.TRY_CATCH_COLLAPSE]: TryCatchCollapse,
    [FlowRendererKey.SLOT_COLLAPSE]: Null,

    [FlowRendererKey.DRAG_NODE]: DragNode,
    [FlowRendererKey.DRAG_HIGHLIGHT_ADDER]: DragHighlightAdder,
    [FlowRendererKey.DRAG_BRANCH_HIGHLIGHT_ADDER]: DragHighlightAdder,
    [FlowRendererKey.DRAGGABLE_ADDER]: DraggingAdder,

    [FlowRendererKey.SELECTOR_BOX_POPOVER]: Null,
  };
};
