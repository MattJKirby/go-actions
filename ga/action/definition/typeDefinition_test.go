package definition

import (
	"go-actions/ga/action"
	"go-actions/ga/cr/asserts"
	ta "go-actions/ga/testing/testActions"
	"go-actions/ga/testing/testHelpers/actionModelTestHelpers"
	"reflect"
	"testing"
)

var mockConfig = &actionModelTestHelpers.MockActionModelConfig{MockUid: "abc"}

func TestTypeDefinitionFromRegistration(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()

	expectedTypeName := "ActionValidEmpty"
	expectedTypePath := "go-actions/ga/testing/testActions/testActions.ActionValidEmpty"
	expectedType := reflect.TypeOf(ta.ActionValidEmpty{})
	expectedValue := reflect.ValueOf(&ta.ActionValidEmpty{})
	expectedCtor := reflect.ValueOf(reg.Constructor).Pointer()
	expectedCtorType := reflect.TypeOf(reg.Constructor)
	expectedPropsType := reflect.TypeOf(*reg.DefaultProps)
	expectedPropsValue := reflect.ValueOf(*reg.DefaultProps)
	expectedTriggerValue := false

	defReg := TypeDefinitionFromRegistration(&reg)

	asserts.Equals(t, expectedTypeName, defReg.TypeName)
	asserts.Equals(t, expectedTypePath, defReg.TypePath)
	asserts.Equals(t, expectedType, defReg.ActionType)
	asserts.Equals(t, expectedValue, defReg.ActionValue)
	asserts.Equals(t, expectedCtor, defReg.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defReg.CtorType)
	asserts.Equals(t, expectedPropsType, defReg.PropsType)
	asserts.Equals(t, expectedPropsValue, defReg.PropsValue)
	asserts.Equals(t, expectedTriggerValue, defReg.Trigger)
}

func TestTriggerDefinitionFromRegistration(t * testing.T) {
	reg := ta.GenerateActionTriggerValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	asserts.Equals(t, true, defReg.Trigger)
}

func TestNewDefaultProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	newProps := defReg.NewDefaultProps()
	typeAssertedProps, ok := newProps.(ta.ActionValidProps)

	asserts.Equals(t, true, ok)
	asserts.Equals(t, *reg.DefaultProps, typeAssertedProps)
}

func TestValidatePropsType(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	tests := []struct {
		name      string
		props     action.GoActionProps
		expectErr bool
	}{
		{name: "valid - same props type", props: ta.ActionValidProps{}, expectErr: false},
		{name: "valid - same props type generated", props: defReg.NewDefaultProps(), expectErr: false},
		{name: "valid - same props type populated", props: ta.ActionValidProps{Param1: "asdf"}, expectErr: false},
		{name: "invalid - same props pointer", props: &ta.ActionValidProps{}, expectErr: true},
		{name: "invalid - different props type", props: ta.ActionInvalidNoExecute{}, expectErr: true},
		{name: "invalid - nil", props: nil, expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := defReg.ValidatePropsType(test.props)
			hasErr := err != nil

			asserts.Equals(t, test.expectErr, hasErr)
		})
	}
}

func TestNewConstructorWithValidProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	expectedInstEmpty := action.NewActionInstance("inst", mockConfig)
	expectedInst := action.NewActionInstance("inst", mockConfig)
	expectedAction := reg.Constructor(expectedInst, ta.ActionValidDefaultProps)
	expectedActionTyped := any(expectedAction).(action.GoAction)

	tests := []struct {
		name             string
		props            action.GoActionProps
		expectedInstance *action.ActionInstance
		expectedAction   action.GoAction
		expectErr        bool
	}{
		{name: "valid", props: defReg.NewDefaultProps(), expectedInstance: expectedInst, expectedAction: expectedActionTyped, expectErr: false},
		{name: "invalid - bad props", props: ta.ActionInvalidNoExecute{}, expectedInstance: expectedInstEmpty, expectedAction: nil, expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInst := action.NewActionInstance("inst", mockConfig)
			testCtor := defReg.NewConstructor()
			testAction, err := testCtor(testInst, test.props)
			hasErr := err != nil

			asserts.Equals(t, test.expectErr, hasErr)
			asserts.Equals(t, test.expectedInstance, testInst)
			asserts.Equals(t, test.expectedAction, testAction)
		})
	}
}

