package app

import (
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"

	"testing"
)

func TestRegisterActionAndGet(t *testing.T) {
	app := NewApp("test")
	registration := ta.GenerateActionValidEmptyRegistration()
	RegisterAction(&registration)(app)

	result, _ := GetActionRegistration[ta.ActionValidEmpty, ta.ActionValidEmptyProps]()(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestGetActionSuccessfulNilProps(t *testing.T) {
	app := NewApp("test")
	registration := ta.GenerateActionValidEmptyRegistration()
	RegisterAction(&registration)(app)

	_, err := GetTypedAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)

	asserts.Equals(t, false, err != nil)
}

func TestGetActionWithProps(t *testing.T) {
	app := NewApp("test")
	reg := ta.GenerateActionValidRegistration()
	RegisterAction(&reg)(app)

	props := &ta.ActionValidProps{Param1: "asdf"}
	act, err := GetTypedAction[ta.ActionValid](props)(app)

	asserts.Equals(t, false, err != nil)
	asserts.Equals(t, props.Param1, act.Action.Param1.Value())

}

func TestGetActionFail(t *testing.T) {
	app := NewApp("test")

	_, err := GetTypedAction[ta.ActionValidEmpty, ta.ActionValidEmptyProps](nil)(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
