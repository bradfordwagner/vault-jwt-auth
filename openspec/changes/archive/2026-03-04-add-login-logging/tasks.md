## 1. Create Token Lookup Package

- [x] 1.1 Create `pkg/vault_token_lookup` directory
- [x] 1.2 Create `pkg/vault_token_lookup/lookup.go` with token lookup function
- [x] 1.3 Add function to call vault `/auth/token/lookup-self` API
- [x] 1.4 Add function to format token data as alphabetically ordered YAML
- [x] 1.5 Ensure token ID field is redacted (show "<redacted>")
- [x] 1.6 Add error handling for lookup failures (non-blocking)

## 2. Update Login Functions to Accept Vault Client

- [x] 2.1 Modify `vault_jwt_login.Login` to return vault client alongside token
- [x] 2.2 Modify `vault_approle_login.Login` to return vault client alongside token
- [x] 2.3 Update function signatures to return `(token string, client *vault.Client, err error)`

## 3. Integrate Token Lookup into JWT Login

- [x] 3.1 Import `pkg/vault_token_lookup` in `pkg/vault_jwt_login/login.go`
- [x] 3.2 Add conditional token lookup call after successful authentication (only if Verbose is true)
- [x] 3.3 Pass vault client and token to lookup function
- [x] 3.4 Display YAML output to stdout (not structured logs)

## 4. Integrate Token Lookup into AppRole Login

- [x] 4.1 Import `pkg/vault_token_lookup` in `pkg/vault_approle_login/login.go`
- [x] 4.2 Add conditional token lookup call after successful authentication (only if Verbose is true)
- [x] 4.3 Pass vault client and token to lookup function
- [x] 4.4 Display YAML output to stdout (not structured logs)

## 5. Update Entra Command for New Return Signature

- [x] 5.1 Update `internal/cmds/entra/entra.go` to handle new return signature from Login
- [x] 5.2 Extract token from new return values
- [x] 5.3 Ensure verbose flag is still passed to Login function

## 6. Error Logging (Already Complete)

- [x] 6.1 Add error logging to JWT login with vault configuration context
- [x] 6.2 Add error logging to AppRole login with vault configuration context
- [x] 6.3 Log vault address, auth endpoint, and role on authentication failures
- [x] 6.4 Update proposal, design, and specs with error logging requirements

## 7. Verbose Flag (Already Complete)

- [x] 7.1 Add `Verbose bool` field to `internal/args/Entra` struct
- [x] 7.2 Find entra cobra command definition and add `-v`/`--verbose` flag
- [x] 7.3 Bind verbose flag to `args.Entra.Verbose` field
- [x] 7.4 Update `internal/cmds/entra/entra.go` to pass `args.Verbose` to `vault_jwt_login.Login`

## 8. Remove Old Logging Code

- [x] 8.1 Remove `pkg/vault_login_logger` directory and files
- [x] 8.2 Remove old logging imports from JWT login
- [x] 8.3 Remove old logging imports from AppRole login

## 9. Testing and Verification

- [x] 9.1 Test JWT login with `-v` flag to verify YAML output appears
- [x] 9.2 Test JWT login without `-v` flag to verify no token lookup occurs
- [x] 9.3 Test AppRole login with verbose to verify YAML output appears
- [x] 9.4 Verify token ID is redacted in YAML output
- [x] 9.5 Verify YAML fields are alphabetically ordered
- [x] 9.6 Verify authentication still succeeds with and without verbose flag
- [x] 9.7 Test authentication failure to verify error logging shows vault context
- [x] 9.8 Verify token lookup failure does not break authentication
