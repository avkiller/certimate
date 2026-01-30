package domain

const CollectionNameWorkflowOutput = "workflow_output"

type WorkflowOutput struct {
	Meta
	WorkflowId string                 `db:"workflowRef" json:"workflowId"`
	RunId      string                 `db:"runRef"      json:"runId"`
	NodeId     string                 `db:"nodeId"      json:"nodeId"`
	NodeConfig WorkflowNodeConfig     `db:"nodeConfig"  json:"nodeConfig"`
	Outputs    []*WorkflowOutputEntry `db:"outputs"     json:"outputs"`
	Succeeded  bool                   `db:"succeeded"   json:"succeeded"`
}

type WorkflowOutputEntry struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	ValueType string `json:"valueType"`
}
