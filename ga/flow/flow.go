package flow

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
)

type Flow struct {
	actions map[string]executable.Action[action.GoAction]
}

func NewFlow() *Flow {
	return &Flow{
		actions: make(map[string]executable.Action[action.GoAction]),
	}
}

func NewAction[T action.GoAction](act *executable.Action[T]) func(*Flow) *executable.Action[T] {
	return func(f *Flow) *executable.Action[T] {
		a, ok := any(act).(executable.Action[action.GoAction])
		if ok {
			fmt.Errorf("could not add action to flow")
		}
		f.actions[act.Instance.Model.ActionUid] = a
		return act
	}
}
