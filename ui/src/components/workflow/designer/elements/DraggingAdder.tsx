import { FlowDragLayer, type AdderProps as FlowgramAdderProps, usePlayground } from "@flowgram.ai/fixed-layout-editor";
import { IconChevronsDown } from "@tabler/icons-react";

export interface DraggingAdderProps extends FlowgramAdderProps {}

const DraggingAdder = ({ from }: DraggingAdderProps) => {
  const playground = usePlayground();

  const layer = playground.getLayer(FlowDragLayer);
  if (!layer) return <></>;
  if (
    layer.options.canDrop &&
    !layer.options.canDrop({
      dragNodes: layer.dragEntities ?? [],
      dropNode: from,
      isBranch: false,
    })
  ) {
    return <></>;
  }

  return (
    <div className="size-4 animate-bounce rounded-full bg-primary text-white shadow-sm">
      <div className="flex size-full items-center justify-center">
        <IconChevronsDown size="1em" />
      </div>
    </div>
  );
};

export default DraggingAdder;
