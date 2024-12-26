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

func NewAction[T action.GoAction](a T) func(*Flow) *executable.Action[T] {
	return func(f *Flow) *executable.Action[T] {
		a, err := app.NewAction[T](a)(f.flowApp)
		if err != nil {
			panic("could not retireve action from app")
		}

		f.actions[a.Instance.Model.ActionUid] = a.Instance
		return a
	}
}
