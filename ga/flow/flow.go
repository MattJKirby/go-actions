package flow

import (
	"encoding/json"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp         *app.App
	actionInstances map[string]*action.ActionInstance
}

type marshalledFlow struct {
	Instances []action.ActionInstance `json:"actions"`
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:         app,
		actionInstances: make(map[string]*action.ActionInstance),
	}
}

func AddAction[T action.GoAction, P action.GoActionProps](props *P) func(*Flow) *executable.Action[T, P] {
	return func(f *Flow) *executable.Action[T, P] {
		act, err := app.GetAction[T](props)(f.flowApp)
		if err != nil {
			panic("could not retrieve action from app")
		}

		f.actionInstances[act.Instance.Model.ActionUid] = act.Instance
		return act
	}
}

func (f *Flow) MarshalJSON() ([]byte, error) {
	Instances := make([]action.ActionInstance, 0)
	for _, instance := range f.actionInstances {
		Instances = append(Instances, *instance)
	}

	return json.Marshal(marshalledFlow{Instances})
}
