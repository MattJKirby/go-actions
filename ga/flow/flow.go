package flow

import (
	"fmt"
	"go-actions/ga/action"
)

type Flow struct {
	actions map[string]action.Action[action.GoAction]
}

func NewFlow() *Flow {
	return &Flow{
		actions: make(map[string]action.Action[action.GoAction]),
	}
}

func NewAction[T action.GoAction](act *action.Action[T]) func (*Flow) *action.Action[T] {
	return func(f *Flow) *action.Action[T] {
		a, ok := any(act).(action.Action[action.GoAction])
		if ok {
			fmt.Errorf("could not add action to flow")
		}
		f.actions[act.Instance.ActionUid] = a
		return act
	}
}