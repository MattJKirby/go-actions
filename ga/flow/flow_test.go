package flow

import (
	"go-actions/ga/app"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"testing"
)

func TestAddAction(t *testing.T) {
	a := app.NewApp("test")
	flow := NewFlow(a)
	app.RegisterAction(testActions.ActionValidEmpty{}, nil)(a)

	act, err := AddAction(flow, func(a testActions.ActionValidEmpty) {})


	assert.Equals(t, nil, err)
	assert.Equals(t, true, act != nil)
	assert.Equals(t, 1, len(flow.Definition.Actions.Store.GetEntries()))
}
