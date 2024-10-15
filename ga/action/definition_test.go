package action

import (
	"go-actions/ga/cr/asserts"
	"go-actions/ga/utils"
	"reflect"
	"testing"
)

type myAction struct {}
func (ma myAction) Execute(){}

func newMyAction() *myAction {
	return &myAction{}
}

func TestNewDef(t *testing.T){
	expectedType := reflect.TypeOf(myAction{})
	expectedValue := reflect.ValueOf(&myAction{})
	expectedCtor := reflect.ValueOf(newMyAction).Pointer()
	expectedName := "myAction"
	expectedTypeName := utils.TypePath(myAction{})
	
	def := NewActionDefinition(newMyAction)
	defType := def.ActionType()
	defValue := def.ActionValue()
	defCtor := def.Constructor().Pointer()
	defName := def.Name()
	defTypeName := def.TypeName()

	t.Run("test def attrs", func(t *testing.T) {
		asserts.Equals(t, defType, expectedType)
		asserts.Equals(t, defValue, expectedValue)
		asserts.Equals(t, expectedCtor, defCtor)
		asserts.Equals(t, expectedName, defName)
		asserts.Equals(t, expectedTypeName, defTypeName)
	})
}