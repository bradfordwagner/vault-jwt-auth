package vault_token_lookup

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/bradfordwagner/go-util/log"
	"github.com/hashicorp/vault-client-go"
	"gopkg.in/yaml.v3"
)

// LookupAndDisplay performs a token lookup and displays the result as alphabetically ordered YAML
func LookupAndDisplay(ctx context.Context, client *vault.Client, token string) error {
	// Set the token on the client for lookup
	if err := client.SetToken(token); err != nil {
		log.Log().With("error", err.Error()).Error("failed to set token on client")
		return err
	}

	// Perform token lookup
	response, err := client.Read(ctx, "/auth/token/lookup-self")
	if err != nil {
		log.Log().With("error", err.Error()).Error("token lookup failed (non-blocking)")
		return err
	}

	if response == nil || response.Data == nil {
		log.Log().Warn("token lookup returned no data")
		return fmt.Errorf("no data returned from token lookup")
	}

	// Extract data and prepare for YAML output
	data := response.Data
	
	// Redact the token ID if present
	if id, ok := data["id"]; ok && id != nil {
		data["id"] = "<redacted>"
	}

	// Convert to YAML with alphabetical ordering
	yamlOutput, err := formatAsOrderedYAML(data)
	if err != nil {
		log.Log().With("error", err.Error()).Error("failed to format token data as YAML")
		return err
	}

	// Output to stdout
	fmt.Println("\n=== Token Information ===")
	fmt.Println(yamlOutput)
	fmt.Println("========================\n")

	return nil
}

// formatAsOrderedYAML converts a map to alphabetically ordered YAML
func formatAsOrderedYAML(data map[string]interface{}) (string, error) {
	// Get sorted keys
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build ordered map
	orderedData := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		orderedData = append(orderedData, map[string]interface{}{key: data[key]})
	}

	// Marshal to YAML
	yamlBytes, err := yaml.Marshal(orderedData)
	if err != nil {
		return "", err
	}

	// Clean up YAML format - remove list markers and extra indentation
	lines := strings.Split(string(yamlBytes), "\n")
	var cleanedLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		// Remove leading "- " from list items
		cleaned := strings.TrimPrefix(line, "- ")
		cleanedLines = append(cleanedLines, cleaned)
	}

	return strings.Join(cleanedLines, "\n"), nil
}
