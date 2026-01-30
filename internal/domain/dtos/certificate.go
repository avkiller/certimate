package dtos

type CertificateDownloadReq struct {
	CertificateId     string `json:"-"`
	CertificateFormat string `json:"format"`
}

type CertificateDownloadResp struct {
	FileBytes  []byte `json:"fileBytes"`
	FileFormat string `json:"fileFormat"`
}

type CertificateRevokeReq struct {
	CertificateId string `json:"-"`
}

type CertificateRevokeResp struct{}
