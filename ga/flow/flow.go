package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app"
)

type Flow struct {
	flowApp *app.App
	actions map[string]*action.ActionInstance
}

// type marshalledFlow struct {
// 	Actions []action.ActionInstance `json:"actions"`
// }

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp: app,
		actions: make(map[string]*action.ActionInstance),
	}
}

func AddAction[T action.GoAction](a T) func(*Flow) *executable.Action[T] {
	return func(f *Flow) *executable.Action[T] {
		act, err := app.NewAction[T](a)(f.flowApp)
		if err != nil {
			panic("could not retireve action from app")
		}

		f.actions[act.Instance.Model.ActionUid] = act.Instance
		return act
	}
}
