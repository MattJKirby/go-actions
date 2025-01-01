package registration

import (
	"go-actions/ga/action"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	th "go-actions/ga/utils/testHelpers"
	"reflect"
	"testing"
)

func TestAcceptAction(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid]()
	registry := NewActionRegistry()
	registration := &action.GoActionRegistration[th.ActionValid]{Constructor: ctor}
	RegisteredAction, err := NewRegisteredAction(registration)

	AcceptAction(RegisteredAction)(registry)
	abt := len(registry.actionsByType)
	abn := len(registry.actionsByName)

	if abt != 1 {
		t.Errorf("test actions by type: got: %d, expected: %d", abt, 1)
	}

	if abn != 1 {
		t.Errorf("test actions by name: got: %d, expected: %d", abt, 1)
	}

	asserts.Equals(t, nil, err)
}

func TestGetAction(t *testing.T) {
	ctor := th.GetEmptyConstructor[th.ActionValid]()
	registration := &action.GoActionRegistration[th.ActionValid]{Constructor: ctor}
	acn, _ := NewRegisteredAction(registration)
	registry := NewActionRegistry()
	AcceptAction(acn)(registry)

	tests := []cr.TestCase[reflect.Type, *RegisteredAction[th.ActionValid]]{
		{Name: "existing def", Input: acn.ActionDefinition.ActionType, Expected: acn},
		{Name: "not existing def", Input: reflect.TypeOf("err"), Expected: nil, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[reflect.Type, *RegisteredAction[th.ActionValid]]) {
		storedDef, err := GetAction[th.ActionValid](test.Input)(registry)

		if test.Error && err == nil {
			t.Errorf("test %s: expected an error but got none", test.Name)
			return
		}

		if !test.Error && storedDef != test.Expected {
			t.Errorf("test %s: got %v, expected %v", test.Name, storedDef, test.Expected)
		}
	})
}
