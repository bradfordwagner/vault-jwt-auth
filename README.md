# vault-jwt-auth
- Used to login to a given provider then auth against vault

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
