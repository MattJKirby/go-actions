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
	reg := ta.GenerateEmptyActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)

	action := NewAction(*def)

	if action.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}

func TestApplyConstructorNoProps(t *testing.T) {
	reg := ta.GenerateEmptyActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)
	acn := NewAction(*def)

	instance := acn.Instance
	action.Parameter(instance, "param", 10)

	newPopulatedInstance(*def, def.DefaultProps)

	asserts.Equals(t, instance, acn.Instance)
}

func TestApplyConstructorWithProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)
	acn := NewAction(*def)

	props := &ta.ActionValidProps{Param1: "AAAA"}
	acn.PopulateActionInstance(props)

	asserts.Equals(t, "AAAA", acn.Action.Param1.Value())

	fmt.Println(acn.Instance)

}
