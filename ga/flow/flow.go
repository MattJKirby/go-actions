package flow

import "go-actions/ga/action"

type Flow struct {
	actions map[string]action.Action
}

func NewFlow() *Flow {
	return &Flow{
		actions: make(map[string]action.Action),
	}
}