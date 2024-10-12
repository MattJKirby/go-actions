package app

import (
	"go-actions/ga/action"
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
	def := action.NewActionDefinition(myActionCtor)
	app.RegisterActionDef(def)

	result, _ := app.GetActionDef(myAction{})

	t.Run("Test get def", func(t *testing.T) {
		if result != def {
			t.Errorf("Error getting definition: expected %v, got %v", &def, result)
		}
	})
}

type myAction struct{}
func (ma myAction) Execute() {}
func myActionCtor() *myAction {
	return &myAction{}
}

func TestNewActionSuccessful(t *testing.T) {
	app := NewApp()
	def := action.NewActionDefinition(myActionCtor)
	app.RegisterActionDef(def)

	t.Run("test new action successful", func(t *testing.T) {
		_, err := NewAction[myAction](myAction{})(app)
		if err != nil {
			t.Errorf("error instatiating action: got %v", nil)
		}
	})
}

func TestNewActionFail(t *testing.T) {
	app := NewApp()

	t.Run("test new action successful", func(t *testing.T) {
		_, err := NewAction[myAction](myAction{})(app)
		if err == nil {
			t.Errorf("error instatiating action: got %v", nil)
		}
	})
}
