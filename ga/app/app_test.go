package app

import (
	"go-actions/ga/action/internals"
	"testing"
)

type myAction struct{}

func (ma myAction) Execute() {}
func myActionCtor(internals.GoActionInternals) *myAction {
	return &myAction{}
}

func TestRegisterActionAndGet(t *testing.T) {
	app := NewApp()
	expected := RegisterAction[myAction](myActionCtor)(app)
	result, _ := app.GetActionDef(myAction{})

	t.Run("Test register def", func(t *testing.T) {
		if result != expected {
			t.Errorf("Error during registration: expected %v, got %v", expected, result)
		}
	})
}

func TestNewActionSuccessful(t *testing.T) {
	app := NewApp()
	RegisterAction(myActionCtor)(app)

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
