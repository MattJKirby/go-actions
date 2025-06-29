package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"testing"
)

func TestAddAction(t *testing.T) {
	a := app.NewApp("test")
	flow := NewFlow(a)
	app.RegisterAction(testActions.ActionValidEmpty{}, nil)(a)
	
	act, err := AddAction(flow, func(a *action.Action[testActions.ActionValidEmpty]) {})
	got, _ := flow.Definition.Actions.Store.Get(act.Instance.Uid.FullUid())

	assert.Equals(t, nil, err)
	assert.Equals(t, true, act != nil)
	assert.Equals(t, true, got != action.ActionInstance{})
}