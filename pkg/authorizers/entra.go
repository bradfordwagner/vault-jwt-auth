package authorizers

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/bradfordwagner/go-util/log"
	"go.uber.org/zap"
)

type Entra struct {
	args EntraArgs
	l    *zap.SugaredLogger
}

type EntraArgs struct {
	ClientId     string
	ClientSecret string
	TenantId     string
}

func NewEntra(args EntraArgs) Authorizer {
	return &Entra{
		l:    log.Log().With("component", "entra_authorizer"),
		args: args,
	}
}

// Authorize implements the Authorizer interface for Entra.
func (e *Entra) Authorize(ctx context.Context) (jwt string, err error) {
	credential, err := azidentity.NewClientSecretCredential(e.args.TenantId, e.args.ClientId, e.args.ClientSecret, &azidentity.ClientSecretCredentialOptions{})
	if err != nil {
		return "", err
	}
	// in order to figure out scopes i ran "az account get-access-token --debug"
	token, err := credential.GetToken(ctx, policy.TokenRequestOptions{
		EnableCAE: true,
		Scopes:    []string{"https://management.core.windows.net//.default"}, // yes the '//' on .default matters
	})
	if err != nil {
		return "", err
	}
	jwt = token.Token
	return
}
