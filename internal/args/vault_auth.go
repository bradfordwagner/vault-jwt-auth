package args

type Vault struct {
	AuthEndpoint string `mapstructure:"VAULT_AUTH_ENDPOINT" env:"VAULT_AUTH_ENDPOINT"`
	AuthRole     string `mapstructure:"VAULT_AUTH_ROLE" env:"VAULT_AUTH_ROLE"`
}
