package main

import (
	"github.com/bradfordwagner/vault-jwt-auth/internal/args"
	"github.com/bradfordwagner/vault-jwt-auth/internal/cmds/entra"

	"github.com/bradfordwagner/go-util/flag_helper"
	"github.com/spf13/cobra"
)

func init() {
	fs := entraVerb.Flags()
	flag_helper.CreateFlag(fs, &entraArgs.Timeout, "timeout", "T", "30s", "Timeout for auth operation (env=TIMEOUT)")
	flag_helper.CreateFlag(fs, &entraArgs.ClientId, "arm_client_id", "i", "", "sp to id to log into entra with (env=ARM_CLIENT_ID)")
	flag_helper.CreateFlag(fs, &entraArgs.ClientSecret, "arm_client_secret", "s", "", "sp secret to log into entra with (env=ARM_CLIENT_SECRET)")
	flag_helper.CreateFlag(fs, &entraArgs.TenantId, "arm_tenant_id", "t", "", "tenant id to log into entra with (env=ARM_TENANT_ID)")
	flag_helper.CreateFlag(fs, &entraVaultArgs.AuthEndpoint, "vault_auth_endpoint", "a", "", "vault to log into entra with (env=VAULT_AUTH_ENDPOINT)")
	flag_helper.CreateFlag(fs, &entraVaultArgs.AuthRole, "vault_auth_role", "r", "", "vault role to log into entra with (env=VAULT_AUTH_ROLE)")
	flag_helper.CreateFlag(fs, &entraArgs.Verbose, "verbose", "v", false, "enable verbose logging for authentication")
	// output
	flag_helper.CreateFlag(fs, &outputArgs.Method, "output_method", "o", "azuredevops", "output method to use for logging (env=OUTPUT_METHOD)")
	flag_helper.CreateFlag(fs, &outputArgs.AzureDevops.Variable, "azure_devops_variable", "", "VAULT_TOKEN", "variable to set in azure devops (env=AZURE_DEVOPS_VARIABLE)")
}

var entraArgs args.Entra
var entraVaultArgs args.Vault
var outputArgs args.Output

var entraVerb = &cobra.Command{
	Use:   "entra",
	Short: "logs microsoft entra",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		flag_helper.Load(&entraArgs)
		flag_helper.Load(&entraVaultArgs)
		flag_helper.Load(&outputArgs)
		entraArgs.Vault = entraVaultArgs
		entraArgs.Output = outputArgs
		return entra.Run(entraArgs)
	},
}
