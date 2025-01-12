package definition

import (
	ta "go-actions/ga/testing/testActions"

	"go-actions/ga/cr/asserts"
	"reflect"
	"testing"
)

func TestActionTypeDefinitionValues(t *testing.T) {
	ctor := ta.GenerateActionValidEmptyCtor()

	expectedType := reflect.TypeOf(ta.ActionValidEmpty{})
	expectedValue := reflect.ValueOf(&ta.ActionValidEmpty{})
	expectedCtor := reflect.ValueOf(ctor).Pointer()
	expectedCtorType := reflect.TypeOf(ctor)
	defCtor, _ := NewTypeDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps](ctor)

	asserts.Equals(t, expectedType, defCtor.ActionType)
	asserts.Equals(t, expectedValue, defCtor.ActionValue)
	asserts.Equals(t, expectedCtor, defCtor.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defCtor.CtorType)
}

func TestConstructorStructParity(t *testing.T) {
	defCtor := TypeDefinitionFromConstructor(ta.GenerateActionValidEmptyCtor())
	defStruct := TypeDefinitionFromStruct[ta.ActionValidEmpty, ta.ActionValidEmptyProps](ta.ActionValidEmpty{})

	asserts.Equals(t, defCtor.ActionType, defStruct.ActionType)
	asserts.Equals(t, defCtor.ActionValue, defStruct.ActionValue)
	asserts.Equals(t, defCtor.CtorType, defStruct.CtorType)
}

func TestNewActionTypeDefiniton(t *testing.T) {
	defCtor, ctorErr := NewTypeDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps](ta.GenerateActionValidEmptyCtor())
	defStruct, structErr := NewTypeDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps](ta.ActionValidEmpty{})
	defInvalid, invalidErr := NewTypeDefinition[ta.ActionValidEmpty, ta.ActionValidEmptyProps](ta.ActionInvalidNoExecute{})

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
