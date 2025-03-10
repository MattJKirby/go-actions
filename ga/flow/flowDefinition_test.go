package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testHelpers/actionModelTestHelpers"
	"testing"
)

func TestAddInstance(t *testing.T){
	flowDef := NewFlowDefinition()
	instance := action.NewActionInstance("someInstance", &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"})

	flowDef.AddInstance(instance)

	asserts.Equals(t, instance, flowDef.actions["someInstance:abc"])
}