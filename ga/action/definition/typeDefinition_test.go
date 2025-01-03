package definition

import (
	th "go-actions/ga/utils/testHelpers"

	"go-actions/ga/cr/asserts"
	"reflect"
	"testing"
)

func TestActionTypeDefinitionValues(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid, th.ActionValidProps]()

	expectedType := reflect.TypeOf(th.ActionValid{})
	expectedValue := reflect.ValueOf(&th.ActionValid{})
	expectedCtor := reflect.ValueOf(ctor).Pointer()
	expectedCtorType := reflect.TypeOf(ctor)
	defCtor, _ := NewTypeDefinition[th.ActionValid, th.ActionValidProps](ctor)

	asserts.Equals(t, expectedType, defCtor.ActionType)
	asserts.Equals(t, expectedValue, defCtor.ActionValue)
	asserts.Equals(t, expectedCtor, defCtor.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defCtor.CtorType)
}

func TestConstructorStructParity(t *testing.T) {
	defCtor := TypeDefinitionFromConstructor(th.GetEmptyConstructor[th.ActionValid, th.ActionValidProps]())
	defStruct := TypeDefinitionFromStruct[th.ActionValid, th.ActionValidProps](th.ActionValid{})

	asserts.Equals(t, defCtor.ActionType, defStruct.ActionType)
	asserts.Equals(t, defCtor.ActionValue, defStruct.ActionValue)
	asserts.Equals(t, defCtor.CtorType, defStruct.CtorType)
}

func TestNewActionTypeDefiniton(t *testing.T) {
	defCtor, ctorErr := NewTypeDefinition[th.ActionValid, th.ActionValidProps](th.GetEmptyConstructor[th.ActionValid, th.ActionValidProps]())
	defStruct, structErr := NewTypeDefinition[th.ActionValid, th.ActionValidProps](th.ActionValid{})
	defInvalid, invalidErr := NewTypeDefinition[th.ActionValid, th.ActionValidProps](th.ActionNoExecute{})

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
