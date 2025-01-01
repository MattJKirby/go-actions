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
  registration := &action.GoActionRegistration[myAction]{Constructor: myActionCtor}
	expected := RegisterAction(registration)(app)

	result, _ := GetActionRegistration[myAction](myAction{})(app)

	if result != expected {
		t.Errorf("Error during registration: expected %v, got %v", expected, result)
	}
}

func TestGetActionSuccessful(t *testing.T) {
	app := NewApp()
  registration := &action.GoActionRegistration[myAction]{Constructor: myActionCtor}
	RegisterAction(registration)(app)

	_, err := GetAction[myAction](myAction{})(app)
	if err != nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}

func TestGetActionFail(t *testing.T) {
	app := NewApp()

	_, err := GetAction[myAction](myAction{})(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
