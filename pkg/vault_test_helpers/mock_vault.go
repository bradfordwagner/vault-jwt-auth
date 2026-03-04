package vault_test_helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault-client-go"
)

// MockVaultClient is a mock implementation of vault client for testing
type MockVaultClient struct {
	WriteFunc     func(ctx context.Context, path string, data map[string]interface{}) (*vault.Response[map[string]interface{}], error)
	ReadFunc      func(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	SetTokenFunc  func(token string) error
	CurrentToken  string
	WriteCallCount int
	ReadCallCount  int
}

// Write mocks the vault client Write method
func (m *MockVaultClient) Write(ctx context.Context, path string, data map[string]interface{}) (*vault.Response[map[string]interface{}], error) {
	m.WriteCallCount++
	if m.WriteFunc != nil {
		return m.WriteFunc(ctx, path, data)
	}
	return nil, fmt.Errorf("WriteFunc not implemented")
}

// Read mocks the vault client Read method
func (m *MockVaultClient) Read(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	m.ReadCallCount++
	if m.ReadFunc != nil {
		return m.ReadFunc(ctx, path, options...)
	}
	return nil, fmt.Errorf("ReadFunc not implemented")
}

// SetToken mocks the vault client SetToken method
func (m *MockVaultClient) SetToken(token string) error {
	m.CurrentToken = token
	if m.SetTokenFunc != nil {
		return m.SetTokenFunc(token)
	}
	return nil
}

// NewMockSuccessfulAuth returns a mock that simulates successful authentication
func NewMockSuccessfulAuth() *MockVaultClient {
	return &MockVaultClient{
		WriteFunc: func(ctx context.Context, path string, data map[string]interface{}) (*vault.Response[map[string]interface{}], error) {
			return &vault.Response[map[string]interface{}]{
				Data: map[string]interface{}{},
				Auth: &vault.ResponseAuth{
					ClientToken:   "test-token-12345678",
					Accessor:      "test-accessor",
					Policies:      []string{"default", "test-policy"},
					TokenPolicies: []string{"default", "test-policy"},
					Metadata: map[string]string{
						"role": "test-role",
					},
					LeaseDuration: 3600,
					Renewable:     true,
					EntityID:      "test-entity-id",
					Orphan:        true,
				},
			}, nil
		},
	}
}

// NewMockFailedAuth returns a mock that simulates failed authentication
func NewMockFailedAuth() *MockVaultClient {
	return &MockVaultClient{
		WriteFunc: func(ctx context.Context, path string, data map[string]interface{}) (*vault.Response[map[string]interface{}], error) {
			return nil, fmt.Errorf("permission denied")
		},
	}
}

// NewMockTokenLookup returns a mock that simulates successful token lookup
func NewMockTokenLookup() *MockVaultClient {
	return &MockVaultClient{
		ReadFunc: func(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
			if path == "/auth/token/lookup-self" {
				return &vault.Response[map[string]interface{}]{
					Data: map[string]interface{}{
						"accessor":         "test-accessor",
						"creation_time":    "2026-03-04T11:47:38Z",
						"creation_ttl":     3600,
						"display_name":     "jwt-test-role",
						"entity_id":        "test-entity-id",
						"expire_time":      "2026-03-04T12:47:38Z",
						"explicit_max_ttl": 0,
						"id":               "test-token-12345678",
						"issue_time":       "2026-03-04T11:47:38Z",
						"meta": map[string]interface{}{
							"role": "test-role",
						},
						"num_uses":  0,
						"orphan":    true,
						"path":      "auth/jwt/test",
						"policies":  []interface{}{"default", "test-policy"},
						"renewable": true,
						"ttl":       3600,
						"type":      "service",
					},
				}, nil
			}
			return nil, fmt.Errorf("path not found")
		},
	}
}

// NewMockFailedTokenLookup returns a mock that simulates failed token lookup
func NewMockFailedTokenLookup() *MockVaultClient {
	return &MockVaultClient{
		ReadFunc: func(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
			return nil, fmt.Errorf("token lookup failed")
		},
	}
}

// NewMockNilResponseTokenLookup returns a mock that returns nil response
func NewMockNilResponseTokenLookup() *MockVaultClient {
	return &MockVaultClient{
		ReadFunc: func(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
			return nil, nil
		},
	}
}
