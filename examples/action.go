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
	ga.DefineAction(&action.GoActionRegistration[ExampleActionConstruct, ExampleActionProps]{
		Constructor: new,
    Props: ExampleActionProps{},
  })
}

type ExampleActionProps struct {
  
}

type ExampleActionConstruct struct {
	IntegerParameter *parameter.ActionParameter[int]
	StringParameter  *parameter.ActionParameter[string]
	Input            *io.Input
	Output           *io.Output
}

func new(instance *action.ActionInstance) *ExampleActionConstruct {
	return &ExampleActionConstruct{
		IntegerParameter: model.Parameter("intParam", 10)(instance.Model),
		StringParameter:  model.Parameter("strParam", "test")(instance.Model),
		Input:            model.Input("input", true)(instance.Model),
		Output:           model.Output("output")(instance.Model),
	}
}

func (ex ExampleActionConstruct) Execute() {
	fmt.Printf("executing Example Action: %d:%s\n", ex.IntegerParameter.Value(), ex.StringParameter.Value())
}
