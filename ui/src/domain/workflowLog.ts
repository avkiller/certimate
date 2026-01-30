export interface WorkflowLogModel extends Omit<BaseModel, "updated"> {
  nodeId: string;
  nodeName: string;
  timestamp: ReturnType<typeof Date.prototype.getTime>;
  level: number;
  message: string;
  data: Record<string, any>;
}

export enum WorkflowLogLevel {
  Debug = -4,
  Info = 0,
  Warn = 4,
  Error = 8,
}
