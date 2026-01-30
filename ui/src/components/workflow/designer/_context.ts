import { createContext, useContext } from "react";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";

export type DesignerContextType = {
  onDocumentChange: () => void;
  onNodeChange: (node: FlowNodeEntity) => void;
  onNodeClick: (node: FlowNodeEntity) => void;
};

export const DesignerContext = createContext<DesignerContextType>({
  onDocumentChange: () => {},
  onNodeChange: () => {},
  onNodeClick: () => {},
});

export const DegisnerContextProvider = DesignerContext.Provider;

export const useDesignerContext = () => {
  const context = useContext(DesignerContext);
  if (!context) {
    throw new Error("`DesignerContext` must be used within a `DesignerContextProvider`");
  }
  return context;
};
