package definition

import (
	"go-actions/ga/action"

	"go-actions/ga/cr/asserts"
	"reflect"
	"testing"
)

type myAction struct{}
type myActionProps struct{}
func (ma myAction) Execute() {}

var newMyAction action.GoActionConstructor[myAction, myActionProps] = func(*action.ActionInstance) *myAction {
	return &myAction{}
}

type invalidAction struct{}

func TestActionTypeDefinitionValues(t *testing.T) {
	expectedType := reflect.TypeOf(myAction{})
	expectedValue := reflect.ValueOf(&myAction{})
	expectedCtor := reflect.ValueOf(newMyAction).Pointer()
	expectedCtorType := reflect.TypeOf(newMyAction)
	defCtor, _ := NewTypeDefinition[myAction, myActionProps](newMyAction)

	asserts.Equals(t, expectedType, defCtor.ActionType)
	asserts.Equals(t, expectedValue, defCtor.ActionValue)
	asserts.Equals(t, expectedCtor, defCtor.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defCtor.CtorType)
}

func TestConstructorStructParity(t *testing.T) {
	defCtor := TypeDefinitionFromConstructor(newMyAction)
	defStruct := TypeDefinitionFromStruct[myAction, myActionProps](myAction{})

	asserts.Equals(t, defCtor.ActionType, defStruct.ActionType)
	asserts.Equals(t, defCtor.ActionValue, defStruct.ActionValue)
	asserts.Equals(t, defCtor.CtorType, defStruct.CtorType)
}

func TestNewActionTypeDefiniton(t *testing.T) {
	defCtor, ctorErr := NewTypeDefinition[myAction, myActionProps](newMyAction)
	defStruct, structErr := NewTypeDefinition[myAction, myActionProps](myAction{})
	defInvalid, invalidErr := NewTypeDefinition[myAction, myActionProps](invalidAction{})

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
