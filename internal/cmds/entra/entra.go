package entra

import (
	"context"
	"github.com/bradfordwagner/go-util/log"
	"github.com/bradfordwagner/vault-jwt-auth/internal/args"
	"github.com/bradfordwagner/vault-jwt-auth/pkg/authorizers"
	"github.com/bradfordwagner/vault-jwt-auth/pkg/output"
	"github.com/bradfordwagner/vault-jwt-auth/pkg/vault_jwt_login"
	"time"
)

func Run(args args.Entra) (err error) {
	l := log.Log()
	l = l.With("client_id", args.ClientId).
		With("tenant_id", args.TenantId)

	// timeout context to limit the login operation
	duration, err := time.ParseDuration(args.Timeout)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), duration)

	// pickup jwt from entra
	auth := authorizers.NewEntra(authorizers.EntraArgs{
		ClientId:     args.ClientId,
		ClientSecret: args.ClientSecret,
		TenantId:     args.TenantId,
	})
	jwt, err := auth.Authorize(ctx)
	if err != nil {
		l.With("error", err.Error()).Error("entra login failed")
		return err
	}
	l.Info("entra login successful")

	// jwt login to vault
	token, _, err := vault_jwt_login.Login(ctx, vault_jwt_login.Args{
		AuthEndpoint: args.Vault.AuthEndpoint,
		Role:         args.Vault.AuthRole,
		Jwt:          jwt,
		Verbose:      args.Verbose,
	})
	if err != nil {
		return err
	}
	l.Info("logged in to vault")

	// setup output method
	o := output.NewOutput(args.Output)
	return o.Output(token)
}
