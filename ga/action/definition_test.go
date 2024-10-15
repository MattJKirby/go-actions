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

	t.Run("test def attrs", func(t *testing.T) {
		asserts.Equals(t, expectedType, def.ActionType())
		asserts.Equals(t, expectedValue, def.ActionValue())
		asserts.Equals(t, expectedCtor, def.Constructor().Pointer())
		asserts.Equals(t, expectedName, def.Name())
		asserts.Equals(t, expectedTypeName, def.TypeName())
	})
}