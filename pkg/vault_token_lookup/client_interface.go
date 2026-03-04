package vault_token_lookup

import (
	"context"

	"github.com/hashicorp/vault-client-go"
)

// VaultClient interface defines the methods we need from vault.Client
// This allows for easier testing with mock implementations
type VaultClient interface {
	Read(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	SetToken(token string) error
}

// Ensure *vault.Client implements VaultClient interface
var _ VaultClient = (*vault.Client)(nil)
