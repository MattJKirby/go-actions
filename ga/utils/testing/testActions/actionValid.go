package testActions

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/parameter"
)

type ActionValid struct {
	Param1 *parameter.ActionParameter[string]
}

type ActionValidProps struct {
	Param1 string
}

var ActionValidDefaultProps = ActionValidProps{Param1: "DefaultParam1Value"}

func (tav ActionValid) Init(inst *action.ActionInstance) {
	tav.Param1 = model.Parameter(inst.Model, "param1", "some val")
}

func (tav ActionValid) Execute() {
	fmt.Println(tav.Param1.Value())
}

func GenerateActionValidRegistration() action.GoActionRegistration[ActionValid] {
	return GenerateRegistration(ActionValid{})
}
