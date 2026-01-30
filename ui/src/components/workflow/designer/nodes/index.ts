import { BizApplyNodeRegistry } from "./BizApplyNodeRegistry";
import { BizDeployNodeRegistry } from "./BizDeployNodeRegistry";
import { BizMonitorNodeRegistry } from "./BizMonitorNodeRegistry";
import { BizNotifyNodeRegistry } from "./BizNotifyNodeRegistry";
import { BizUploadNodeRegistry } from "./BizUploadNodeRegistry";
import { BranchBlockNodeRegistry, ConditionNodeRegistry } from "./ConditionNode";
import { DelayNodeRegistry } from "./DelayNode";
import { EndNodeRegistry } from "./EndNode";
import { StartNodeRegistry } from "./StartNode";
import { CatchBlockNodeRegistry, TryCatchNodeRegistry } from "./TryCatchNode";

export const getAllNodeRegistries = () => {
  return [
    StartNodeRegistry,
    EndNodeRegistry,
    DelayNodeRegistry,
    BizApplyNodeRegistry,
    BizUploadNodeRegistry,
    BizMonitorNodeRegistry,
    BizDeployNodeRegistry,
    BizNotifyNodeRegistry,
    ConditionNodeRegistry,
    BranchBlockNodeRegistry,
    TryCatchNodeRegistry,
    CatchBlockNodeRegistry,
  ];
};

export type * from "./typings";
