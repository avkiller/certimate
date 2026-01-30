package aliyunapigw

const (
	// 服务类型：原 API 网关。
	SERVICE_TYPE_TRADITIONAL = "traditional"
	// 服务类型：云原生 API 网关。
	SERVICE_TYPE_CLOUDNATIVE = "cloudnative"
)

const (
	// 匹配模式：精确匹配。
	DOMAIN_MATCH_PATTERN_EXACT = "exact"
	// 匹配模式：通配符匹配。
	DOMAIN_MATCH_PATTERN_WILDCARD = "wildcard"
	// 匹配模式：证书 SAN 匹配。
	DOMAIN_MATCH_PATTERN_CERTSAN = "certsan"
)
