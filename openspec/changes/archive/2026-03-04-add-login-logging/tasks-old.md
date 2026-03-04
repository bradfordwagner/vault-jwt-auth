## 1. Create Shared Logging Package

- [x] 1.1 Create `pkg/vault_login_logger` directory
- [x] 1.2 Create `pkg/vault_login_logger/logger.go` with common logging function
- [x] 1.3 Add function to extract and log auth info from vault response with nil checks
- [x] 1.4 Implement role_id redaction (first 8 chars only)
- [x] 1.5 Ensure no sensitive data (tokens, secrets) are logged

## 2. Add Verbose Field to Login Args

- [x] 2.1 Add `Verbose bool` field to `vault_jwt_login.Args` struct
- [x] 2.2 Add `Verbose bool` field to `vault_approle_login.Args` struct

## 3. Integrate Logging into JWT Login

- [x] 3.1 Import `github.com/bradfordwagner/go-util/log` in `pkg/vault_jwt_login/login.go`
- [x] 3.2 Import `pkg/vault_login_logger` in `pkg/vault_jwt_login/login.go`
- [x] 3.3 Add conditional logging call after successful `vaultClient.Write()` (only if Verbose is true)
- [x] 3.4 Pass auth endpoint and role to logging function

## 4. Integrate Logging into AppRole Login

- [x] 4.1 Import `github.com/bradfordwagner/go-util/log` in `pkg/vault_approle_login/login.go`
- [x] 4.2 Import `pkg/vault_login_logger` in `pkg/vault_approle_login/login.go`
- [x] 4.3 Add conditional logging call after successful `vaultClient.Write()` (only if Verbose is true)
- [x] 4.4 Pass auth endpoint and role_id to logging function

## 5. Add Verbose Flag to Entra Command

- [x] 5.1 Add `Verbose bool` field to `internal/args/Entra` struct
- [x] 5.2 Find entra cobra command definition and add `-v`/`--verbose` flag
- [x] 5.3 Bind verbose flag to `args.Entra.Verbose` field
- [x] 5.4 Update `internal/cmds/entra/entra.go` to pass `args.Verbose` to `vault_jwt_login.Login`

## 6. Add Error Logging

- [x] 6.1 Add error logging to JWT login with vault configuration context
- [x] 6.2 Add error logging to AppRole login with vault configuration context
- [x] 6.3 Log vault address, auth endpoint, and role on authentication failures
- [x] 6.4 Update proposal, design, and specs with error logging requirements

## 7. Testing and Verification

- [ ] 7.1 Test JWT login with `-v` flag to verify logging output appears
- [ ] 7.2 Test JWT login without `-v` flag to verify no logging output
- [ ] 7.3 Test AppRole login with verbose to verify logging output appears
- [ ] 7.4 Verify no sensitive data (JWT, secret_id, client_token) in logs
- [ ] 7.5 Verify authentication still succeeds with and without verbose flag
- [ ] 7.6 Verify logging handles missing auth metadata gracefully
- [ ] 7.7 Test authentication failure to verify error logging shows vault context
