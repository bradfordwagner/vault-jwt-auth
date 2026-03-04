# Agent Configuration

## Context Loading

When working on this project, always load and reference the OpenAPI specification located in the `openspec/` directory.

### OpenAPI Specification
- **Location**: `openspec/config.yaml`
- **Purpose**: This project uses a spec-driven development approach. The OpenAPI specification defines the API contracts and should be consulted when making changes to API endpoints, request/response structures, or authentication flows.

## Project Context

This is a Vault JWT authentication tool (`vault-jwt-auth`) that:
- Logs into identity providers (e.g., Microsoft Entra)
- Authenticates against HashiCorp Vault using JWT tokens
- Outputs tokens for use in CI/CD pipelines (Azure DevOps)

## Development Guidelines

1. Always check the OpenAPI spec before modifying API-related code
2. Follow the spec-driven schema defined in `openspec/config.yaml`
3. Maintain consistency with existing authentication flows
4. Consider both CLI and container-based usage patterns

### OpenSpec Workflow (CRITICAL)

**ALWAYS follow the OpenSpec workflow for ALL changes:**

1. **Propose** (`/opsx-propose`) - Create proposal, design, specs, and tasks BEFORE coding
2. **Update Specs** - For ANY code change, update the corresponding spec files FIRST
   - `openspec/changes/<change-name>/specs/**/*.md` - Requirements and scenarios
   - `openspec/changes/<change-name>/design.md` - Technical decisions
   - `openspec/changes/<change-name>/proposal.md` - What and why
3. **Apply** (`/opsx-apply`) - Implement tasks from the spec
4. **Archive** (`/opsx-archive`) - Archive completed changes

**Never skip updating specs when making code changes.** Specs are the source of truth.

## Build and Testing

- **Build Command**: Use `go install ./cmd/vja` instead of `go build`
  - This installs the binary to `$GOPATH/bin` so `vja` can be run directly from anywhere
  - Easier for integration testing and local development
