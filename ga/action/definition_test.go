package action

import (
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
		if defType != expectedType{
			t.Errorf("expected %v, got %v",expectedType, defType)
		}

		if defValue != expectedValue {
			t.Errorf("expected %v, got %v",expectedValue, defValue)
		}

		if expectedCtor != defCtor {
			t.Errorf("expected %v, got %v",expectedCtor, defCtor)
		}

		if expectedName != defName {
			t.Errorf("expected %v, got %v",expectedName, defName)
		}

		if expectedTypeName != defTypeName {
			t.Errorf("expected %v, got %v",expectedTypeName, defTypeName)
		}
	})
}