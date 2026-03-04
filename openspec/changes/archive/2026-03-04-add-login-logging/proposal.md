## Why

Both JWT and AppRole login methods currently lack visibility into authentication details. Adding structured logging will enable debugging, auditing, and monitoring of vault authentication flows across both methods.

## What Changes

- Add vault token lookup functionality after successful authentication
- Integrate bradfordwagner logger (`github.com/bradfordwagner/go-util/log`) into both login packages
- Log token metadata from vault token lookup (not from login response) when verbose flag is enabled
- Output token information in multi-line, alphabetically ordered YAML format for readability
- Log vault configuration context (address, endpoint, role) on authentication failures
- Ensure sensitive data (JWT tokens, secret_ids, client tokens) are never logged or displayed
- Add verbose flag (`-v`, `--verbose`) to entra command to enable detailed token lookup logging

## Capabilities

### New Capabilities
- `token-lookup-logging`: Perform vault token lookup after authentication and log token metadata in YAML format for both JWT and AppRole methods
- `verbose-flag`: Add verbose flag to entra command to control token lookup logging output

### Modified Capabilities
<!-- No existing capabilities are being modified -->

## Impact

- **Code**: `pkg/vault_jwt_login/login.go` and `pkg/vault_approle_login/login.go` will be modified to add logging; `internal/cmds/entra/entra.go` will be modified to accept verbose flag
- **CLI**: New optional `-v`/`--verbose` flag for entra command (backward compatible)
- **Dependencies**: Already using `github.com/bradfordwagner/go-util/log` in the project (via `go.mod`)
- **APIs**: No breaking changes - logging is additive and doesn't change function signatures
- **Systems**: Improved observability for vault authentication flows
