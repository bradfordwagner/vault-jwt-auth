## ADDED Requirements

### Requirement: Perform vault token lookup after authentication
The system SHALL perform a vault token lookup using the `/auth/token/lookup-self` API after successful authentication.

#### Scenario: Token lookup after JWT login
- **WHEN** JWT login succeeds and verbose flag is enabled
- **THEN** system calls vault token lookup API with the authenticated token

#### Scenario: Token lookup after AppRole login
- **WHEN** AppRole login succeeds and verbose flag is enabled
- **THEN** system calls vault token lookup API with the authenticated token

#### Scenario: Token lookup skipped when verbose disabled
- **WHEN** authentication succeeds and verbose flag is false
- **THEN** system does not perform token lookup

### Requirement: Display token information in YAML format
The system SHALL output token information in multi-line, alphabetically ordered YAML format.

#### Scenario: YAML output is multi-line
- **WHEN** token information is displayed
- **THEN** system outputs each field on a separate line

#### Scenario: YAML fields are alphabetically ordered
- **WHEN** token information is displayed
- **THEN** system orders all fields alphabetically (accessor, creation_time, etc.)

#### Scenario: YAML is valid and parseable
- **WHEN** token information is displayed
- **THEN** system outputs valid YAML that can be parsed by standard YAML parsers

### Requirement: Display all non-sensitive token fields
The system SHALL display all token metadata fields except the actual token value.

#### Scenario: Display standard token fields
- **WHEN** token lookup succeeds
- **THEN** system displays: accessor, creation_time, creation_ttl, display_name, entity_id, expire_time, explicit_max_ttl, issue_time, meta, num_uses, orphan, path, policies, renewable, ttl, type

#### Scenario: Redact token ID
- **WHEN** token lookup response includes token ID
- **THEN** system displays id field as "<redacted>" instead of actual token value

#### Scenario: Handle missing optional fields
- **WHEN** token lookup response lacks optional fields
- **THEN** system displays available fields without failing

### Requirement: Token lookup uses authenticated token
The system SHALL use the newly acquired token to perform the lookup-self operation.

#### Scenario: Token is set on client before lookup
- **WHEN** performing token lookup
- **THEN** system calls client.SetToken() with the authenticated token before calling lookup-self

#### Scenario: Lookup uses current token
- **WHEN** performing token lookup
- **THEN** system uses the token returned from authentication, not a different token

### Requirement: Token lookup failure does not break authentication
The system SHALL continue normal operation if token lookup fails.

#### Scenario: Lookup fails but authentication succeeds
- **WHEN** token lookup API call fails
- **THEN** system logs error but still returns the authentication token successfully

#### Scenario: Lookup timeout does not block
- **WHEN** token lookup times out
- **THEN** system logs timeout error and continues without blocking authentication flow

### Requirement: Verbose flag controls token lookup
The system SHALL only perform token lookup when verbose flag is enabled.

#### Scenario: Verbose enabled triggers lookup
- **WHEN** verbose flag is true and authentication succeeds
- **THEN** system performs token lookup and displays YAML output

#### Scenario: Verbose disabled skips lookup
- **WHEN** verbose flag is false and authentication succeeds
- **THEN** system does not perform token lookup or display token information
