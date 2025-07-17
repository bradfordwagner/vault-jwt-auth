package args

type OutputMethod string

const (
	AzureDevops OutputMethod = "azuredevops"
)

func ToOutputMethod(s string) OutputMethod {
	switch s {
	case "azuredevops":
		return AzureDevops
	default:
		return AzureDevops // default to AzureDevops if unknown
	}
}

type Output struct {
	Method      string `mapstructure:"output_method" env:"OUTPUT_METHOD"`
	AzureDevops AzureDevopsOutput
}

type AzureDevopsOutput struct {
	Variable string `mapstructure:"variable" env:"AZURE_DEVOPS_VARIABLE"`
}
