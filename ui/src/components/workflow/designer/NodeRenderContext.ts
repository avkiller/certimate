import { createContext, useContext } from "react";
import { type NodeRenderReturnType } from "@flowgram.ai/fixed-layout-editor";

export type NodeRenderContextType = NodeRenderReturnType;

export const NodeRenderContext = createContext<NodeRenderContextType>({} as NodeRenderContextType);

export const NodeRenderContextProvider = NodeRenderContext.Provider;

export const useNodeRenderContext = () => {
  const context = useContext(NodeRenderContext);
  if (!context) {
    throw new Error("`NodeRenderContext` must be used within a `NodeRenderContextProvider`");
  }
  return context;
};
