package args

type Entra struct {
	Vault Vault

	Output Output

	ClientId     string `mapstructure:"ARM_CLIENT_ID" env:"ARM_CLIENT_ID"`
	ClientSecret string `mapstructure:"ARM_CLIENT_SECRET" env:"ARM_CLIENT_SECRET"`
	TenantId     string `mapstructure:"ARM_TENANT_ID" env:"ARM_TENANT_ID"`
}
