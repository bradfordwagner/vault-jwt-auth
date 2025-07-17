package authorizers

import (
	"context"
	"template_cli/internal/args"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/bradfordwagner/go-util/log"
	"go.uber.org/zap"
)

type Entra struct {
	args args.Entra
	l    *zap.SugaredLogger
}

func NewEntra(entra args.Entra) Authorizer {
	return &Entra{
		l:    log.Log().With("component", "entra_authorizer"),
		args: entra,
	}
}

// Authorize implements the Authorizer interface for Entra.
func (e *Entra) Authorize() (jwt string, err error) {
	credential, err := azidentity.NewClientSecretCredential(e.args.TenantId, e.args.ClientId, e.args.ClientSecret, &azidentity.ClientSecretCredentialOptions{})
	if err != nil {
		return "", err
	}
	// in order to figure out scopes i ran "az account get-access-token --debug"
	token, err := credential.GetToken(context.Background(), policy.TokenRequestOptions{
		EnableCAE: true,
		Scopes:    []string{"https://management.core.windows.net//.default"}, // yes the '//' on .default matters
	})
	if err != nil {
		return "", err
	}
	jwt = token.Token
	return
}
