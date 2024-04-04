package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/net/context"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &EchoFunction{}

type EchoFunction struct{}

func (f *EchoFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "echo"
}

func (f *EchoFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Echo a string",
		Description: "Given a string value, returns the same value.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "Value to echo",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *EchoFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	// Read Terraform argument data into the variable
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))

	// Set the result to the same data
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, input))
}

func NewEchoFunction() function.Function {
	return &EchoFunction{}
}
