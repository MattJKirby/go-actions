package flow

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers/actionTestHelpers"
	"testing"
)

// func TestAddInstance(t *testing.T) {
// 	flowDef := NewFlowDefinition()
// 	instance := action.NewActionInstance("someInstance", &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"})
// 	instance.Model.ActionTrigger

// 	flowDef.AddInstance(instance)

// 	asserts.Equals(t, instance, flowDef.Actions["someInstance:abc"])
// }

// func TestAddTriggerInstance(t *testing.T) {
// 	flowDef := NewFlowDefinition()
// 	triggerInstance := action.NewActionInstance("triggerInstance", &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"})

// 	flowDef.AddInstance(triggerInstance)

// 	asserts.Equals(t, triggerInstance, flowDef.Triggers["triggerInstance:abc"])
// 	asserts.Equals(t, 0, len(flowDef.Actions))
// }

func TestNewAction(t *testing.T) {
	a := app.NewApp("testApp")
	reg := testActions.GenerateActionValidEmptyRegistration()
	app.RegisterAction(&reg)(a)

	flowDef := NewFlowDefinition()
	act, err := flowDef.NewAction(a, "ActionValidEmpty")

	assert.Equals(t, true, flowDef.Actions[act.InitialisedInstance.Model.ActionUid] != nil)
	assert.Equals(t, nil, err)
}

func TestMarshalFlowDefinition(t *testing.T) {
	flowDef := NewFlowDefinition()
	instance := action.NewActionInstance("someInstance", &actionTestHelpers.MockActionConfig{MockUid: "abc"})

	flowDef.AddInstance(instance)
	marshalledInstance, _ := json.Marshal(instance)

	marshalled, err := json.Marshal(flowDef)
	expected := fmt.Sprintf(`{"Actions":{"%s":%s}}`, instance.Model.ActionUid, marshalledInstance)
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, string(marshalled))
}
