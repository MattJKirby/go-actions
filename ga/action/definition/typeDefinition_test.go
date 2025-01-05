package definition

import (
	ta "go-actions/ga/testing/testActions"

	"go-actions/ga/cr/asserts"
	"reflect"
	"testing"
)

func TestActionTypeDefinitionValues(t *testing.T) {
	ctor := ta.GenerateActionValidCtor()

	expectedType := reflect.TypeOf(ta.ActionValid{})
	expectedValue := reflect.ValueOf(&ta.ActionValid{})
	expectedCtor := reflect.ValueOf(ctor).Pointer()
	expectedCtorType := reflect.TypeOf(ctor)
	defCtor, _ := NewTypeDefinition[ta.ActionValid, ta.ActionValidProps](ctor)

	asserts.Equals(t, expectedType, defCtor.ActionType)
	asserts.Equals(t, expectedValue, defCtor.ActionValue)
	asserts.Equals(t, expectedCtor, defCtor.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defCtor.CtorType)
}

func TestConstructorStructParity(t *testing.T) {
	defCtor := TypeDefinitionFromConstructor(ta.GenerateActionValidCtor())
	defStruct := TypeDefinitionFromStruct[ta.ActionValid, ta.ActionValidProps](ta.ActionValid{})

	asserts.Equals(t, defCtor.ActionType, defStruct.ActionType)
	asserts.Equals(t, defCtor.ActionValue, defStruct.ActionValue)
	asserts.Equals(t, defCtor.CtorType, defStruct.CtorType)
}

func TestNewActionTypeDefiniton(t *testing.T) {
	defCtor, ctorErr := NewTypeDefinition[ta.ActionValid, ta.ActionValidProps](ta.GenerateActionValidCtor())
	defStruct, structErr := NewTypeDefinition[ta.ActionValid, ta.ActionValidProps](ta.ActionValid{})
	defInvalid, invalidErr := NewTypeDefinition[ta.ActionValid, ta.ActionValidProps](ta.ActionInvalidNoExecute{})

	if defCtor == nil || ctorErr != nil {
		t.Errorf("error generating Type Def from constructor: expected %v, got %v", nil, ctorErr)
	}

	if defStruct == nil || structErr != nil {
		t.Errorf("error generating Type Def from struct: expected %v, got %v", nil, structErr)
	}

	if defInvalid != nil || invalidErr == nil {
		t.Errorf("error generating Type Def from invalid: expected err, got %v", nil)
	}
}
