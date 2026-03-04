package vault_jwt_login

import (
	"context"
	"github.com/bradfordwagner/go-util/log"
	"github.com/bradfordwagner/vault-jwt-auth/pkg/vault_token_lookup"
	"github.com/hashicorp/vault-client-go"
)

type Args struct {
	VaultAddress string
	AuthEndpoint string
	Role         string
	Jwt          string
	Verbose      bool
}

func Login(ctx context.Context, a Args) (token string, client *vault.Client, err error) {
	// init vault client
	opts := []vault.ClientOption{
		vault.WithEnvironment(),
	}
	if a.VaultAddress != "" {
		opts = append(opts, vault.WithAddress(a.VaultAddress))
	}
	vaultClient, err := vault.New(opts...)
	if err != nil {
		log.Log().With("error", err.Error()).Error("failed to create vault client")
		return "", nil, err
	}

	//token=$(vault write ${auth_endpoint}/login role=${role} jwt="${jwt}" | qq '.auth.client_token' -r)
	loginPath := a.AuthEndpoint + "/login"
	write, err := vaultClient.Write(ctx, loginPath, map[string]interface{}{
		"role": a.Role,
		"jwt":  a.Jwt,
	})
	if err != nil {
		// Get vault address for logging
		vaultAddr := a.VaultAddress
		if vaultAddr == "" {
			vaultAddr = "(from VAULT_ADDR env)"
		}
		log.Log().With("vault_address", vaultAddr).
			With("auth_endpoint", a.AuthEndpoint).
			With("auth_role", a.Role).
			With("error", err.Error()).
			Error("vault jwt login failed")
		return "", nil, err
	}

	token = write.Auth.ClientToken

	if a.Verbose {
		// Perform token lookup and display YAML output
		_ = vault_token_lookup.LookupAndDisplay(ctx, vaultClient, token)
	}

	return token, vaultClient, nil
}
