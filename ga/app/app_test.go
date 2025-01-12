package app

import (
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"

	"testing"
)

func TestRegisterActionAndGet(t *testing.T) {
	app := NewApp()
	registration := ta.GenerateEmptyActionValidRegistration()
	RegisterAction(&registration)(app)

	result, _ := GetActionRegistration[ta.EmptyActionValid, ta.EmptyActionValidProps]()(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestGetActionSuccessfulNilProps(t *testing.T) {
	app := NewApp()
	registration := ta.GenerateEmptyActionValidRegistration()
	RegisterAction(&registration)(app)

	act, err := GetAction[ta.EmptyActionValid, ta.EmptyActionValidProps](nil)(app)

	asserts.Equals(t, false, err != nil)
	asserts.Equals(t, new(ta.EmptyActionValidProps), act.Props)
}


// TODO
func TestGetActionWithProps(t *testing.T){
	app := NewApp()
	reg := ta.GenerateEmptyActionValidRegistration()
	RegisterAction(&reg)(app)

	// props := &ta.EmptyActionValidProps{Prop: "asdf"}
	// act, err := GetAction[ta.EmptyActionValid](props)(app)
	
	// asserts.Equals(t, false, err != nil)
	// asserts.Equals(t, props, act.Props)

}

func TestGetActionFail(t *testing.T) {
	app := NewApp()

	_, err := GetAction[ta.EmptyActionValid, ta.EmptyActionValidProps](nil)(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
