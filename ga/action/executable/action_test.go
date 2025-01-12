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

	applyConstructor(*def, instance, nil)

	asserts.Equals(t, instance, acn.Instance)
}


//TODO
func TestApplyConstructorWithProps(t *testing.T){
	reg := ta.GenerateEmptyActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)
	acn := NewAction(*def)


	// applyConstructor(*def, acn.Instance, props)

	fmt.Println(acn.Instance)

}