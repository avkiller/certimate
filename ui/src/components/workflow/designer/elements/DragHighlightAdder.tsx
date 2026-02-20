import { type DragNodeProps as FlowgramDragNodeProps } from "@flowgram.ai/fixed-layout-editor";
import { IconGradienter } from "@tabler/icons-react";

const DragHighlightAdder = (_: FlowgramDragNodeProps) => {
  return (
    <div className="size-4 animate-ping rounded-full bg-primary text-white shadow-sm">
      <div className="flex size-full items-center justify-center">
        <IconGradienter size="1em" />
      </div>
    </div>
  );
};

export default DragHighlightAdder;
