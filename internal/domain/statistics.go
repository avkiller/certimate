package domain

type Statistics struct {
	CertificateTotal        int `json:"certificateTotal"`
	CertificateExpiringSoon int `json:"certificateExpiringSoon"`
	CertificateExpired      int `json:"certificateExpired"`

	WorkflowTotal    int `json:"workflowTotal"`
	WorkflowEnabled  int `json:"workflowEnabled"`
	WorkflowDisabled int `json:"workflowDisabled"`
}
