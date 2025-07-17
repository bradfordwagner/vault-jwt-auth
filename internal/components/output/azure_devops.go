package output

import (
	"fmt"
	"github.com/bradfordwagner/go-util/log"
	"go.uber.org/zap"
	"template_cli/internal/args"
)

type azureDevops struct {
	l *zap.SugaredLogger
	a args.Output
}

func NewAzureDevops(a args.Output) Output {
	return &azureDevops{
		l: log.Log().With("output_method", "azuredevops"),
		a: a,
	}
}

func (a *azureDevops) Output(token string) (err error) {
	//https://learn.microsoft.com/en-us/azure/devops/pipelines/process/set-variables-scripts?view=azure-devops&tabs=bash#set-variable-properties
	o := fmt.Sprintf("##vso[task.setvariable variable=%s;issecret=true]%s", a.a.AzureDevops.Variable, token)
	_, err = fmt.Println(o)
	return
}
