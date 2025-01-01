package app

import (
	"go-actions/ga/action"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestRegisterActionAndGet(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid]()
	app := NewApp()
	registration := &action.GoActionRegistration[th.ActionValid]{Constructor: ctor}
	expected := RegisterAction(registration)(app)

	result, _ := GetActionRegistration[th.ActionValid](th.ActionValid{})(app)

	if result != expected {
		t.Errorf("Error during registration: expected %v, got %v", expected, result)
	}
}

func TestGetActionSuccessful(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid]()
	app := NewApp()
	registration := &action.GoActionRegistration[th.ActionValid]{Constructor: ctor}
	RegisterAction(registration)(app)

	_, err := GetAction[th.ActionValid](th.ActionValid{})(app)
	if err != nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}

func TestGetActionFail(t *testing.T) {
	app := NewApp()

	_, err := GetAction[th.ActionValid](th.ActionValid{})(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
