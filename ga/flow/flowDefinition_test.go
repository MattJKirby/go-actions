package flow

import (
	"go-actions/ga/app"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidProvider: mockGenerator}

func TestAddActionDefinition(t *testing.T) { 
	a := app.NewApp("testApp")
	app.RegisterAction[testActions.ActionValidEmpty](nil)(a)

	flowDef := NewFlowDefinition(a)
	typedef, _ := app.GetDefinitionByName("ActionValidEmpty")(a)

	act, err := addAction[testActions.ActionValidEmpty](flowDef, typedef)

	got, _ := flowDef.Actions.GetResource(act.Instance.Uid.FullUid())
	assert.Equals(t, true, got != nil)
	assert.Equals(t, nil, err)
}

func TestNewAction(t *testing.T) {
	a := app.NewApp("testApp")
	app.RegisterAction[testActions.ActionValidEmpty](nil)(a)

	flowDef := NewFlowDefinition(a)
	act, err := flowDef.NewAction("ActionValidEmpty")

	got, _ := flowDef.Actions.GetResource(act.Instance.Uid.FullUid())
	assert.Equals(t, true, got != nil)
	assert.Equals(t, nil, err)
}

// func TestNewReference(t *testing.T) {
// 	a := app.NewApp("testApp")
// }
