package app

import (
	ta "go-actions/ga/testing/testActions"

	"testing"
)

func TestRegisterActionAndGet(t *testing.T) {
	app := NewApp()
	registration := ta.GenerateActionValidRegistration()
	RegisterAction(&registration)(app)

	result, _ := GetActionRegistration[ta.ActionValid, ta.ActionValidProps](ta.ActionValid{})(app)

	if result == nil {
		t.Errorf("Error during registration: expected %v, got %v", nil, result)
	}
}

func TestGetActionSuccessful(t *testing.T) {
	app := NewApp()
	registration := ta.GenerateActionValidRegistration()
	RegisterAction(&registration)(app)

	_, err := GetAction[ta.ActionValid, ta.ActionValidProps](ta.ActionValid{}, nil)(app)
	if err != nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}

func TestGetActionFail(t *testing.T) {
	app := NewApp()

	_, err := GetAction[ta.ActionValid, ta.ActionValidProps](ta.ActionValid{}, nil)(app)
	if err == nil {
		t.Errorf("error instatiating action: got %v", nil)
	}
}
