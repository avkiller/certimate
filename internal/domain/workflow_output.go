package domain

const CollectionNameWorkflowOutput = "workflow_output"

type WorkflowOutput struct {
	Meta
	WorkflowId string                 `json:"workflowId" db:"workflowRef"`
	RunId      string                 `json:"runId" db:"runRef"`
	NodeId     string                 `json:"nodeId" db:"nodeId"`
	NodeConfig WorkflowNodeConfig     `json:"nodeConfig" db:"nodeConfig"`
	Outputs    []*WorkflowOutputEntry `json:"outputs" db:"outputs"`
	Succeeded  bool                   `json:"succeeded" db:"succeeded"`
}

type WorkflowOutputEntry struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	ValueType string `json:"valueType"`
}
