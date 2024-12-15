package examples

import (
	"fmt"
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
)

func init() {
	ga.DefineAction(NewExampleAction)
}

type ExampleAction struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *io.Input
	Output           *io.Output
}

func NewExampleAction(instance action.ActionInstance) *ExampleAction {
	return &ExampleAction{
		IntegerParameter: model.Parameter("intParam", 10)(instance),
		StringParameter:  model.Parameter("strParam", "test")(instance.ActionModel),
		Input:            model.Input("input", true)(instance.ActionModel),
		Output:           model.Output("output")(instance.ActionModel),
	}
}

func (ex ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}
