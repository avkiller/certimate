package kong

type Certificate struct {
	Id        *string   `json:"id,omitempty"`
	Cert      *string   `json:"cert,omitempty"`
	CertAlt   *string   `json:"cert_alt,omitempty"`
	Key       *string   `json:"key,omitempty"`
	KeyAlt    *string   `json:"key_alt,omitempty"`
	SNIs      []*string `json:"snis,omitempty"`
	Tags      []*string `json:"tags,omitempty"`
	CreatedAt *string   `json:"created_at,omitempty"`
	UpdatedAt *string   `json:"updated_at,omitempty"`
}
