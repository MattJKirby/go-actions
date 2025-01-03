package definition

import (
	"go-actions/ga/action"
	"go-actions/ga/cr/asserts"
	th "go-actions/ga/utils/testHelpers"
	"testing"
)

func TestNewDef(t *testing.T) {
	expectedTypeName := "ActionValid"
	expectedTypePath := "go-actions/ga/utils/testHelpers/testHelpers.ActionValid"
	ctor := th.GetEmptyConstructor[th.ActionValid, th.ActionValidProps]()
	reg := action.GoActionRegistration[th.ActionValid, th.ActionValidProps]{Constructor: ctor}

	defCtor, _ := NewActionDefinition(&reg)

	asserts.Equals(t, defCtor.Name, expectedTypeName)
	asserts.Equals(t, defCtor.TypePath, expectedTypePath)
	// asserts.Equals(t, defCtor.ActionType, defCtor.ActionType)
	// asserts.Equals(t, defCtor.CtorValue, defCtor.CtorValue)
}
