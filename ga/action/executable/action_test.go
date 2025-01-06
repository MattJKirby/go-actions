package executable

import (
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"testing"
)

func TestNewAction(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)

	action := NewAction(*def)

	if action.Instance == nil {
		t.Errorf("invalid action: instance expected but got nil")
	}
}

func TestApplyConstructorNoProps(t *testing.T){
	reg := ta.GenerateActionValidRegistration()
	def, _ := definition.NewActionDefinition(&reg)
	action := NewAction(*def)
	
	instance := *action.Instance
	model.Parameter("param", 10)(instance.Model)

	applyConstructor(*def, &instance, nil)

	asserts.Equals(t, instance, *action.Instance)
}
