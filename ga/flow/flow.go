package flow

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp        *app.App
	flowDefinition *flowDefinition
}

func NewFlow(app *app.App, flowDefinition *flowDefinition) *Flow {
	return &Flow{
		flowApp:        app,
		flowDefinition: flowDefinition,
	}
}

func NewFlowAction[T action.GoAction, P action.GoActionProps](f *Flow, props *P) (*executable.BaseExecutable[T], error) {
	instantiated, err := app.GetAction[T](props)(f.flowApp)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve action from app")
	}

	f.flowDefinition.AddInstance(instantiated.Instance)

	test, _ := json.Marshal(instantiated.Instance)
	fmt.Println(string(test))
	return instantiated, nil
}
