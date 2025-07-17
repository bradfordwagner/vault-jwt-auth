package vault_jwt_login

import (
	"context"
	"github.com/hashicorp/vault-client-go"
)

type Args struct {
	VaultAddress string
	AuthEndpoint string
	Role         string
	Jwt          string
}

func Login(ctx context.Context, a Args) (token string, err error) {
	// init vault client
	opts := []vault.ClientOption{
		vault.WithEnvironment(),
	}
	if a.VaultAddress != "" {
		opts = append(opts, vault.WithAddress(a.VaultAddress))
	}
	vaultClient, err := vault.New(opts...)

	//token=$(vault write ${auth_endpoint}/login role=${role} jwt="${jwt}" | qq '.auth.client_token' -r)
	loginPath := a.AuthEndpoint + "/login"
	write, err := vaultClient.Write(ctx, loginPath, map[string]interface{}{
		"role": a.Role,
		"jwt":  a.Jwt,
	})
	if err != nil {
		return
	}
	token = write.Auth.ClientToken
	return
}
