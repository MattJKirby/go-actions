package flow

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testHelpers/actionModelTestHelpers"
	"testing"
)

func TestAddInstance(t *testing.T) {
	flowDef := NewFlowDefinition()
	instance := action.NewActionInstance("someInstance", &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"})

	flowDef.AddInstance(instance)

	asserts.Equals(t, instance, flowDef.Actions["someInstance:abc"])
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
