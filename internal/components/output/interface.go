package output

import "template_cli/internal/args"

type Output interface {
	Output(token string) (err error)
}

func NewOutput(a args.Output) Output {
	switch args.ToOutputMethod(a.Method) {
	case args.AzureDevops:
		return NewAzureDevops(a)
	default:
		return NewAzureDevops(a) // default to AzureDevops if unknown
	}
}
