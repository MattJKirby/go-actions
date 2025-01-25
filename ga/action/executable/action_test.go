package executable

import (
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"

	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestNewAction(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()
	def := definition.NewActionDefinition(&reg)

	action := NewAction(*def, nil)

	if action.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}

func TestApplyConstructorNoProps(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()
	def := definition.NewActionDefinition(&reg)
	acn := NewAction(*def, nil)

	instance := acn.Instance
	action.Parameter(instance, "param", 10)

	newPopulatedInstance(*def, def.DefaultProps)

	asserts.Equals(t, instance, acn.Instance)
}

func TestApplyConstructorWithProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	def := definition.NewActionDefinition(&reg)
	props := &ta.ActionValidProps{Param1: "AAAA"}
	acn := NewAction(*def, props)

	asserts.Equals(t, "AAAA", acn.Action.Param1.Value())

	fmt.Println(acn.Instance)

}
