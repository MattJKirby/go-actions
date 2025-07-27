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
	ga.RegisterAction(&ExampleAction{}, nil)
}

type ExampleAction struct {
	*action.Internals
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *io.ActionInput
	Output           *io.ActionOutput
}

func (ex *ExampleAction) Init(inst *action.ActionInstance) {
	ex.IntegerParameter = model.Parameter(inst.Model, "intParam", 10)
	ex.StringParameter = model.Parameter(inst.Model, "strParam", "str")
	ex.Input = model.Input(inst.Model, "input", true, nil)
	ex.Output = model.Output(inst.Model, "output", nil)
}

func (ex *ExampleAction) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}
