import { createContext, useContext } from "react";
import { type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";

// #region FormNestedFieldsContext
export type FormNestedFieldsContextType = {
  parentNamePath: string;
};

export const FormNestedFieldsContext = createContext<FormNestedFieldsContextType>({
  parentNamePath: "",
});

export const FormNestedFieldsContextProvider = FormNestedFieldsContext.Provider;

export const useFormNestedFieldsContext = () => {
  const context = useContext(FormNestedFieldsContext);
  if (!context) {
    throw new Error("`FormNestedFieldsContext` must be used within a `FormNestedFieldsContextProvider`");
  }
  return context;
};
// #endregion

// #region NodeFormContext
export type NodeFormContextType = {
  node: FlowNodeEntity;
};

export const NodeFormContext = createContext<NodeFormContextType>({} as NodeFormContextType);

export const NodeFormContextProvider = NodeFormContext.Provider;

export const useNodeFormContext = () => {
  const context = useContext(NodeFormContext);
  if (!context) {
    throw new Error("`NodeFormContext` must be used within a `NodeFormContextProvider`");
  }
  return context;
};
// #endregion
