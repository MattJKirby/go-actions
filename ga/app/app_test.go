package app

import (
	"go-actions/ga/action"
	"testing"
)

type myAction struct{}

func (ma myAction) Execute() {}
func myActionCtor(*action.ActionInstance) *myAction {
	return &myAction{}
}

func TestRegisterActionAndGet(t *testing.T) {
	app := NewApp()
	expected := DefineAction[myAction](myActionCtor)(app)
	result, _ := app.GetActionDef(myAction{})

	if result != expected {
		t.Errorf("Error during registration: expected %v, got %v", expected, result)
	}
}

func TestNewActionSuccessful(t *testing.T) {
	app := NewApp()
	DefineAction(myActionCtor)(app)

	_, err := NewAction[myAction](myAction{})(app)
	if err != nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}

func TestNewActionFail(t *testing.T) {
	app := NewApp()

	_, err := NewAction[myAction](myAction{})(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
