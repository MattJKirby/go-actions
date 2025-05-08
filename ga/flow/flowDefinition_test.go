package flow

import (
	"go-actions/ga/action/model"
	"go-actions/ga/app"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidProvider{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidProvider: mockGenerator}

func TestNewAction(t *testing.T) {
	a := app.NewApp("testApp")
	app.RegisterAction[testActions.ActionValidEmpty](nil)(a)

	flowDef := NewFlowDefinition(a)
	act, err := flowDef.NewAction("ActionValidEmpty")

	got, _ := flowDef.Actions.GetResource(act.Instance.Uid.FullUid())
	assert.Equals(t, true, got != nil)
	assert.Equals(t, nil, err)
}

func TestGetModels(t *testing.T) {
	a := app.NewApp("testApp")
	app.RegisterAction[testActions.ActionValidEmpty](nil)(a)

	flowDef := NewFlowDefinition(a)
	a1, _ := flowDef.NewAction("ActionValidEmpty")
	a2, _ := flowDef.NewAction("ActionValidEmpty")

	expected := []*model.ActionModel{
		a1.Instance.Model,
		a2.Instance.Model,
	}

	assert.Equals(t, expected, flowDef.GetModels())
}

// func TestNewReference(t *testing.T) {
// 	a := app.NewApp("testApp")
// }
