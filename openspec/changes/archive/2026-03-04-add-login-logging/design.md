## Context

The vault-jwt-auth project provides two authentication methods for HashiCorp Vault: JWT-based and AppRole-based login. Both methods are implemented in separate packages (`pkg/vault_jwt_login` and `pkg/vault_approle_login`) with similar structure but no shared logging infrastructure. The project already uses `github.com/bradfordwagner/go-util/log` for logging in other parts of the codebase (e.g., `internal/cmds/entra`).

Currently, there's no visibility into authentication flows, making debugging and auditing difficult. The vault client library returns a `*vault.Response` object with auth metadata that could be logged.

## Goals / Non-Goals

**Goals:**
- Add vault token lookup functionality after successful authentication
- Create a reusable mechanism to query and display token information via vault API
- Output token metadata in multi-line, alphabetically ordered YAML format
- Log vault configuration context on authentication failures for debugging
- Ensure sensitive data (tokens, secrets) are never logged or displayed
- Add verbose flag to entra command to control token lookup output

**Non-Goals:**
- Changing the public API of either login function
- Adding configurable log levels beyond verbose on/off
- Logging failed authentication attempts (error handling already exists)
- Performance optimization of login flows
- Adding verbose flag to other commands (only entra for now)

## Decisions

### 1. Token Lookup vs Response Logging
**Decision**: Use vault token lookup API (`/auth/token/lookup-self`) instead of logging from login response.

**Rationale**: Token lookup provides authoritative, consistent token information directly from vault. This approach:
- Gets the same data structure regardless of auth method (JWT, AppRole, etc.)
- Ensures we're showing current token state, not just initial response
- Allows vault to control what information is exposed
- Provides a single code path for all auth methods

**Alternatives Considered**:
- Log from login response: Rejected as different auth methods return different response structures
- Parse response.Auth directly: Rejected as it couples logging to response format changes

### 2. YAML Output Format
**Decision**: Output token information as multi-line, alphabetically ordered YAML.

**Rationale**: YAML format provides:
- Human-readable structure that's easy to scan visually
- Alphabetical ordering makes it easy to find specific fields
- Multi-line format is more readable than single-line JSON
- Familiar format for operations teams
- Easy to copy/paste for documentation or tickets

**Alternatives Considered**:
- JSON output: Rejected as less readable for humans
- Table format: Rejected as harder to copy/paste and less structured
- Single-line structured logging: Rejected as too compact for detailed token info

### 3. What to Display
**Decision**: Display all non-sensitive fields from token lookup response, alphabetically ordered:
- accessor
- creation_time
- creation_ttl
- display_name
- entity_id
- expire_time
- explicit_max_ttl
- id (redacted - only show "<redacted>")
- issue_time
- meta (metadata map)
- num_uses
- orphan
- path
- policies
- renewable
- ttl
- type

**Rationale**: Show all available information except the actual token value. This provides complete operational visibility while maintaining security.

**Alternatives Considered**:
- Show subset of fields: Rejected as users may need different fields for different debugging scenarios
- Show token value: Rejected for security reasons

### 4. Token Lookup Integration Point
**Decision**: Perform token lookup immediately after successful authentication, before returning the token. Set the authenticated token on the vault client before calling lookup-self.

**Rationale**: This ensures we only lookup successful authentications and use the correct token. The vault client must be configured with the new token via `client.SetToken()` before performing the lookup-self operation.

**Alternatives Considered**:
- Lookup in calling code: Rejected as it would require changes to all callers
- Lookup before vault call: Rejected as we wouldn't have a token yet
- Use client without setting token: Rejected as lookup-self requires the token to be set on the client

### 5. Verbose Flag Control
**Decision**: Add a `Verbose bool` field to login Args structs and pass it through from entra command. Only log when verbose is true.

**Rationale**: Login logging may be too noisy for production use. A verbose flag gives users control over when to enable detailed auth logging. The entra command will accept `-v`/`--verbose` flag and pass it to the login functions.

**Alternatives Considered**:
- Always log: Rejected as it may be too verbose for production
- Environment variable: Rejected as CLI flag is more explicit and easier to use
- Log level configuration: Rejected as overkill for this use case

## Risks / Trade-offs

**Risk**: Logging adds latency to authentication flow  
→ **Mitigation**: Logging is synchronous but minimal (structured logging is fast). If this becomes an issue, can switch to async logging later.

**Risk**: Log output may expose information in certain environments  
→ **Mitigation**: Only logging non-sensitive fields. Token accessors are explicitly safe per Vault docs.

**Risk**: Changes to vault-client-go response structure could break logging  
→ **Mitigation**: Use defensive nil checks. Logging failures should not break authentication flow.

## Migration Plan

1. Create new `pkg/vault_login_logger` package with shared logging function
2. Add `Verbose bool` field to `vault_jwt_login.Args` and `vault_approle_login.Args`
3. Update login functions to call logger only when `Verbose` is true
4. Add `Verbose bool` field to `internal/args/Entra` struct
5. Update entra command to accept `-v`/`--verbose` flag and pass to login
6. Test both login methods with and without verbose flag
7. No deployment/rollback concerns - changes are additive and backward compatible (verbose defaults to false)
