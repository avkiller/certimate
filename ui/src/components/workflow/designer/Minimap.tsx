import { useService } from "@flowgram.ai/fixed-layout-editor";
import { FlowMinimapService, MinimapRender } from "@flowgram.ai/minimap-plugin";

export interface MinimapProps {
  className?: string;
  style?: React.CSSProperties;
}

const Minimap = ({ className, style }: MinimapProps) => {
  const minimapService = useService(FlowMinimapService);

  return (
    <div className={className} style={style}>
      <MinimapRender
        service={minimapService}
        panelStyles={{}}
        containerStyles={{
          pointerEvents: "auto",
          position: "relative",
          top: "unset",
          right: "unset",
          bottom: "unset",
          left: "unset",
        }}
        inactiveStyle={{
          opacity: 1,
          scale: 0.5,
          translateX: 0,
          translateY: 0,
        }}
      />
    </div>
  );
};

export default Minimap;
