package entra

import (
	"context"
	"github.com/bradfordwagner/go-util/log"
	"github.com/hashicorp/vault-client-go"
	"template_cli/internal/args"
	"template_cli/internal/components/authorizers"
	"template_cli/internal/components/output"
)

func Run(args args.Entra) (err error) {
	l := log.Log()
	l = l.With("client_id", args.ClientId).
		With("tenant_id", args.TenantId)

	// pickup jwt from entra
	auth := authorizers.NewEntra(args)
	jwt, err := auth.Authorize()
	if err != nil {
		l.With("error", err.Error()).Error("entra login failed")
		return err
	}
	l.Info("entra login successful")

	// if vault auth endpoint is set, use it to login to vault
	vaultClient, err := vault.New(
		vault.WithEnvironment(),
	)
	//token=$(vault write ${auth_endpoint}/login role=${role} jwt="${jwt}" | qq '.auth.client_token' -r)
	loginPath := args.Vault.AuthEndpoint + "/login"
	l = l.
		With("vault_addr", vaultClient.Configuration().Address).
		With("login_path", loginPath).
		With("role", args.Vault.AuthRole)
	write, err := vaultClient.Write(context.TODO(), loginPath, map[string]interface{}{
		"role": args.Vault.AuthRole,
		"jwt":  jwt,
	})
	if err != nil {
		l.With("error", err.Error()).Error("failed to login to vault")
		return err
	}
	l.Info("logged in to vault")
	token := write.Auth.ClientToken

	// setup output method
	o := output.NewOutput(args.Output)
	return o.Output(token)
}
