# vault-jwt-auth
- Used to login to a given provider then auth against vault

## Installation
```bash
go install github.com/bradfordwagner/vault-jwt-auth/cmd/vja@latest
```

## Usage
```bash
logs microsoft entra

Usage:
  vja entra [flags]

Flags:
  -i, --arm_client_id string           sp to id to log into entra with (env=ARM_CLIENT_ID)
  -s, --arm_client_secret string       sp secret to log into entra with (env=ARM_CLIENT_SECRET)
  -t, --arm_tenant_id string           tenant id to log into entra with (env=ARM_TENANT_ID)
      --azure_devops_variable string   variable to set in azure devops (env=AZURE_DEVOPS_VARIABLE) (default "VAULT_TOKEN")
  -h, --help                           help for entra
  -o, --output_method string           output method to use for logging (env=OUTPUT_METHOD) (default "azuredevops")
  -v, --verbose                        enable verbose output (displays token information in YAML format)
  -a, --vault_auth_endpoint string     vault to log into entra with (env=VAULT_AUTH_ENDPOINT)
  -r, --vault_auth_role string         vault role to log into entra with (env=VAULT_AUTH_ROLE)
```

## Verbose Mode

Enable verbose mode to see detailed token information after successful authentication:

```bash
vja entra -v
```

This will display token metadata in alphabetically ordered YAML format, including:
- Token accessor
- Policies
- TTL and expiration times
- Entity ID
- Metadata
- And more (token value is redacted for security)

Example output:
```yaml
=== Token Information ===
accessor: hmac-sha256:a1b2c3d4...
creation_time: 2026-03-04T11:47:38Z
creation_ttl: 3600
display_name: jwt-my_role
entity_id: c30886d4-df27-f9b0-b8c6-7b9c4c3b4701
expire_time: 2026-03-04T12:47:38Z
id: <redacted>
policies:
  - default
  - read_secrets
renewable: true
ttl: 3600
========================
```

## From a docker container
```bash
docker run -it --rm \
  -e VAULT_AUTH_ENDPOINT \
  -e VAULT_AUTH_ROLE \
  -e ARM_CLIENT_ID \
  -e ARM_CLIENT_SECRET \
  -e ARM_TENANT_ID \
  -e VAULT_ADDR \
  ghcr.io/bradfordwagner/vault-jwt-auth:0.1.0-rc2 entra
```

## Development
```bash
cat <<EOF > .env
export VAULT_AUTH_ENDPOINT=auth/jwt/tenant
export VAULT_AUTH_ROLE=...
export ARM_CLIENT_ID=...
export ARM_CLIENT_SECRET=...
export ARM_TENANT_ID=...
export VAULT_ADDR=...
EOF
chmod 700 .env

# Load environment variables
. ./env
```
