package config

type Zitadel struct {
	Issuer  string `koanf:"issuer"`
	ApiUrl  string `koanf:"api_url"`
	OrgId   string `koanf:"org_id"`
	JwtPath string `koanf:"jwt_path"`
}
