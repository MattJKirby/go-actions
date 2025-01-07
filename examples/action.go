package examples

import (
	"fmt"
	"go-actions/ga"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/flow"
)

func init() {
	ga.RegisterAction(&action.GoActionRegistration[ExampleAction, ExampleActionProps]{
		Constructor:  newExampleAction,
		DefaultProps: new(ExampleActionProps),
	})
}

type ExampleAction struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *io.Input
	Output           *io.Output
}

type ExampleActionProps struct{}

func newExampleAction(instance *action.ActionInstance, props ExampleActionProps) *ExampleAction {
	return &ExampleAction{
		IntegerParameter: model.Parameter("intParam", 10)(instance.Model),
		StringParameter:  model.Parameter("strParam", "test")(instance.Model),
		Input:            model.Input("input", true)(instance.Model),
		Output:           model.Output("output")(instance.Model),
	}
}

func (ex ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}

func NewExampleAction(flow *flow.Flow, props *ExampleActionProps) *executable.Action[ExampleAction, ExampleActionProps] {
	return ga.NewFlowAction(flow, ExampleAction{}, props)
}
