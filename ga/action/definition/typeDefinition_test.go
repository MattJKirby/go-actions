package definition

import (
	th "go-actions/ga/utils/testHelpers"

	"go-actions/ga/cr/asserts"
	"reflect"
	"testing"
)

func TestActionTypeDefinitionValues(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.TestActionValid]()
	
	expectedType := reflect.TypeOf(th.TestActionValid{})
	expectedValue := reflect.ValueOf(&th.TestActionValid{})
	expectedCtor := reflect.ValueOf(ctor).Pointer()
	expectedCtorType := reflect.TypeOf(ctor)
	defCtor, _ := NewTypeDefinition[th.TestActionValid](ctor)

	asserts.Equals(t, expectedType, defCtor.ActionType)
	asserts.Equals(t, expectedValue, defCtor.ActionValue)
	asserts.Equals(t, expectedCtor, defCtor.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defCtor.CtorType)
}

func TestConstructorStructParity(t *testing.T) {
	defCtor := TypeDefinitionFromConstructor(th.GetEmptyConstructor[th.TestActionValid]())
	defStruct := TypeDefinitionFromStruct(th.TestActionValid{})

	asserts.Equals(t, defCtor.ActionType, defStruct.ActionType)
	asserts.Equals(t, defCtor.ActionValue, defStruct.ActionValue)
	asserts.Equals(t, defCtor.CtorType, defStruct.CtorType)
}

func TestNewActionTypeDefiniton(t *testing.T) {
	defCtor, ctorErr := NewTypeDefinition[th.TestActionValid](th.GetEmptyConstructor[th.TestActionValid]())
	defStruct, structErr := NewTypeDefinition[th.TestActionValid](th.TestActionValid{})
	defInvalid, invalidErr := NewTypeDefinition[th.TestActionValid](th.TestActionNoExecute{})

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
