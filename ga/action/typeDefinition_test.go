package action

import (
	"go-actions/ga/cr/asserts"
	"reflect"
	"testing"
)

type myAction struct {}
func (ma myAction) Execute(){}
var newMyAction Constructor[myAction] = func() *myAction {
	return &myAction{}
}

type invalidAction struct{}

func TestActionTypeDefinitionValues(t *testing.T){
	expectedType := reflect.TypeOf(myAction{})
	expectedValue := reflect.ValueOf(&myAction{})
	expectedCtor := reflect.ValueOf(newMyAction).Pointer()
	expectedCtorType := reflect.TypeOf(newMyAction)
	defCtor, _ := NewTypeDefinition[myAction](newMyAction)

	t.Run("test def attrs from constructor", func(t *testing.T) {
		asserts.Equals(t, expectedType, defCtor.actionType)
		asserts.Equals(t, expectedValue, defCtor.actionValue)
		asserts.Equals(t, expectedCtor, defCtor.ctorValue.Pointer())
		asserts.Equals(t, expectedCtorType, defCtor.ctorType)
	})
}

func TestConstructorStructParity(t *testing.T){
	defCtor := TypeDefinitionFromConstructor(newMyAction)
	defStruct := TypeDefinitionFromStruct(myAction{})

	t.Run("test def attrs from constructor", func(t *testing.T) {
		asserts.Equals(t, defCtor.actionType, defStruct.actionType)
		asserts.Equals(t, defCtor.actionValue, defStruct.actionValue)
		asserts.Equals(t, defCtor.ctorType, defStruct.ctorType)
	})
}

func TestNewActionTypeDefiniton(t *testing.T){
	defCtor, ctorErr := NewTypeDefinition[myAction](newMyAction)
	defStruct, structErr := NewTypeDefinition[myAction](myAction{})
	defInvalid, invalidErr := NewTypeDefinition[myAction](invalidAction{})

	t.Run("assert correct type", func(t *testing.T) {
		if defCtor == nil || ctorErr != nil {
			t.Errorf("error generating Type Def from constructor: expected %v, got %v", nil, ctorErr)
		}

		if defStruct == nil || structErr != nil {
			t.Errorf("error generating Type Def from struct: expected %v, got %v", nil, structErr)
		}

		if defInvalid != nil || invalidErr == nil {
			t.Errorf("error generating Type Def from invalid: expected err, got %v", nil)
		}
	})
}
