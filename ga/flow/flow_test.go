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
	app.RegisterAction[testActions.ActionValidEmpty](nil)(a)
	
	act, err := AddAction[testActions.ActionValidEmpty](flow)

	assert.Equals(t, nil, err)
	assert.Equals(t, true, act != nil)
}