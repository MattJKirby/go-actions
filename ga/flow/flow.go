package flow

import (
	"fmt"
	"go-actions/ga/action"
)

type Flow struct {
	actions map[string]action.GoAction[action.Action]
}

func NewFlow() *Flow {
	return &Flow{
		actions: make(map[string]action.GoAction[action.Action]),
	}
}

func NewAction[T action.Action](act *action.GoAction[T]) func (*Flow) *action.GoAction[T] {
	return func(f *Flow) *action.GoAction[T] {
		a, ok := any(act).(action.GoAction[action.Action])
		if ok {
			fmt.Errorf("could not add action to flow")
		}
		f.actions[act.Instance.ActionUid] = a
		return act
	}
}