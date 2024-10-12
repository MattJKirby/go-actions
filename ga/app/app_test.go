package app

import (
	"go-actions/ga/action"
	"reflect"
	"testing"
)

func TestRegisterDef(t *testing.T) {
	app := NewApp()
	def := action.ActionDefinition{}

	result := app.RegisterActionDef(&def)

	t.Run("Test register def", func(t *testing.T) {
		if result != &def {
			t.Errorf("Error during registration: expected %v, got %v", &def, result)
		}
	})
}

func TestGetActionDef(t *testing.T) {
	app := NewApp()
	def := action.ActionDefinition{}
	app.RegisterActionDef(&def)

	result, _ := app.GetActionDef(def.ActionType())

	t.Run("Test get def", func(t *testing.T) {
		if result != &def {
			t.Errorf("Error getting definition: expected %v, got %v", &def, result)
		}
	})
}

type myRegisteredAction struct{}
func (ma myRegisteredAction) Execute() {}
func myActionCtor() *myRegisteredAction {
	return &myRegisteredAction{}
}

func TestNewActionSuccessful(t *testing.T) {
	app := NewApp()
	def := action.NewActionDefinition(myActionCtor)
	app.RegisterActionDef(def)

	t.Run("test new action successful", func(t *testing.T) {
		_, err := NewAction[myRegisteredAction](myRegisteredAction{})(app)
		if err != nil {
			t.Errorf("error instatiating action: got %v", nil)
		}
	})
}

func TestNewActionFail(t *testing.T) {
	app := NewApp()

	t.Run("test new action successful", func(t *testing.T) {
		_, err := NewAction[myRegisteredAction](myRegisteredAction{})(app)
		if err == nil {
			t.Errorf("error instatiating action: got %v", nil)
		}
	})
}


func TestGetActionTypeFromDefType(t *testing.T){
	app := NewApp()
	expected := reflect.TypeOf(myRegisteredAction{})
	actionType := app.getActionFromType(myRegisteredAction{})


	t.Run("test action from type retrieval", func(t *testing.T) {
		if actionType != expected {
			t.Errorf("error getting action type: expected %v but got %v", expected, actionType)
		}
	})
}
