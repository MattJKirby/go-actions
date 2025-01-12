package app

import (
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"

	"testing"
)

func TestRegisterActionAndGet(t *testing.T) {
	app := NewApp()
	registration := ta.GenerateActionValidRegistration()
	RegisterAction(&registration)(app)

	result, _ := GetActionRegistration[ta.ActionValid, ta.ActionValidProps]()(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestGetActionSuccessfulNilProps(t *testing.T) {
	app := NewApp()
	registration := ta.GenerateActionValidRegistration()
	RegisterAction(&registration)(app)

	act, err := GetAction[ta.ActionValid, ta.ActionValidProps](nil)(app)

	asserts.Equals(t, false, err != nil)
	asserts.Equals(t, new(ta.ActionValidProps), act.Props)
}

func TestGetActionWithProps(t *testing.T){
	app := NewApp()
	reg := ta.GenerateActionValidRegistration()
	RegisterAction(&reg)(app)

	props := &ta.ActionValidProps{Prop: "asdf"}
	act, err := GetAction[ta.ActionValid](props)(app)
	
	asserts.Equals(t, false, err != nil)
	asserts.Equals(t, props, act.Props)

}

func TestGetActionFail(t *testing.T) {
	app := NewApp()

	_, err := GetAction[ta.ActionValid, ta.ActionValidProps](nil)(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
