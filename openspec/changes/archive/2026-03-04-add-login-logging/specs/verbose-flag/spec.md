## ADDED Requirements

### Requirement: Entra command accepts verbose flag
The system SHALL accept a verbose flag (`-v` or `--verbose`) in the entra command.

#### Scenario: User provides short verbose flag
- **WHEN** user runs entra command with `-v` flag
- **THEN** system enables verbose logging for authentication

#### Scenario: User provides long verbose flag
- **WHEN** user runs entra command with `--verbose` flag
- **THEN** system enables verbose logging for authentication

#### Scenario: User omits verbose flag
- **WHEN** user runs entra command without verbose flag
- **THEN** system runs with verbose disabled (default behavior)

### Requirement: Verbose flag controls token lookup
The system SHALL only perform token lookup and display when verbose flag is enabled.

#### Scenario: Verbose enabled triggers token lookup
- **WHEN** verbose flag is true and login succeeds
- **THEN** system performs vault token lookup and displays YAML output

#### Scenario: Verbose disabled skips token lookup
- **WHEN** verbose flag is false and login succeeds
- **THEN** system does not perform token lookup or display token information

### Requirement: Verbose flag is passed to login functions
The system SHALL pass verbose flag value from entra command to login functions for token lookup control.

#### Scenario: Verbose flag propagates to JWT login
- **WHEN** entra command receives verbose flag
- **THEN** system passes verbose value to vault_jwt_login.Login via Args struct

#### Scenario: Verbose flag propagates to AppRole login
- **WHEN** approle command receives verbose flag
- **THEN** system passes verbose value to vault_approle_login.Login via Args struct

### Requirement: Login Args structs support verbose field
The system SHALL add Verbose field to JWT and AppRole login Args structs.

#### Scenario: JWT Args has Verbose field
- **WHEN** JWT login Args is constructed
- **THEN** system provides Verbose bool field

#### Scenario: AppRole Args has Verbose field
- **WHEN** AppRole login Args is constructed
- **THEN** system provides Verbose bool field

### Requirement: Backward compatibility maintained
The system SHALL maintain backward compatibility when verbose flag is not provided.

#### Scenario: Existing usage without verbose works
- **WHEN** entra command is called without verbose flag
- **THEN** system behaves exactly as before (no logging, authentication succeeds)
