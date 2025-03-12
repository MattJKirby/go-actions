package flow

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testActions"
	"go-actions/ga/testing/testHelpers/actionModelTestHelpers"
	"testing"
)

func TestAddInstance(t *testing.T) {
	flowDef := NewFlowDefinition()
	instance := action.NewActionInstance("someInstance", &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"})

	flowDef.AddInstance(instance)

	asserts.Equals(t, instance, flowDef.Actions["someInstance:abc"])
}

func TestNewAction(t *testing.T) {
	a := app.NewApp("testApp")
	reg := testActions.GenerateActionValidEmptyRegistration()
	app.RegisterAction(&reg)(a)

	flowDef := NewFlowDefinition()
	act, err := flowDef.NewAction(a, "ActionValidEmpty")

	asserts.Equals(t, true, flowDef.Actions[act.InitialisedInstance.Model.ActionUid] != nil)
	asserts.Equals(t, nil, err)
}

func TestMarshalFlowDefinition(t *testing.T) {
	flowDef := NewFlowDefinition()
	instance := action.NewActionInstance("someInstance", &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"})

	flowDef.AddInstance(instance)
	marshalledInstance, _ := json.Marshal(instance)

	marshalled, err := json.Marshal(flowDef)
	expected := fmt.Sprintf(`{"Actions":{"%s":%s}}`, instance.Model.ActionUid, marshalledInstance)
	asserts.Equals(t, nil, err)
	asserts.Equals(t, expected, string(marshalled))
}
