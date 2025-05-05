package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestNewAction(t *testing.T) {
	a := app.NewApp("testApp")
	reg := action.ActionRegistration[testActions.ActionValidEmpty]{Action: testActions.ActionValidEmpty{}}
	app.RegisterAction(&reg)(a)

	flowDef := NewFlowDefinition(a)
	act, err := flowDef.NewAction("ActionValidEmpty")

	got, _ := flowDef.Actions.GetResource(act.Instance.Uid.FullUid())
	assert.Equals(t, true, got != nil)
	assert.Equals(t, nil, err)
}

// func TestNewReference(t *testing.T) {
// 	a := app.NewApp("testApp")
// }
