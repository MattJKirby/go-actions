package definition

import (
	"go-actions/ga/action"
	"go-actions/ga/testing/assert"
	ta "go-actions/ga/testing/testActions"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"
	"reflect"
	"testing"
)

var mockConfig = &actionTestHelpers.MockActionConfig{MockUid: "abc"}

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

	assert.Equals(t, expectedTypeName, defReg.TypeName)
	assert.Equals(t, expectedTypePath, defReg.TypePath)
	assert.Equals(t, expectedType, defReg.ActionType)
	assert.Equals(t, expectedValue, defReg.ActionValue)
	assert.Equals(t, expectedCtor, defReg.CtorValue.Pointer())
	assert.Equals(t, expectedCtorType, defReg.CtorType)
	assert.Equals(t, expectedPropsType, defReg.PropsType)
	assert.Equals(t, expectedPropsValue, defReg.PropsValue)
	assert.Equals(t, expectedTriggerValue, defReg.Trigger)
}

func TestTriggerDefinitionFromRegistration(t *testing.T) {
	reg := ta.GenerateActionTriggerValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	assert.Equals(t, true, defReg.Trigger)
}

func TestNewDefaultProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	newProps := defReg.NewDefaultProps()
	typeAssertedProps, ok := newProps.(ta.ActionValidProps)

	assert.Equals(t, true, ok)
	assert.Equals(t, *reg.DefaultProps, typeAssertedProps)
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

			assert.Equals(t, test.expectErr, hasErr)
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

			assert.Equals(t, test.expectErr, hasErr)
			assert.Equals(t, test.expectedInstance, testInst)
			assert.Equals(t, test.expectedAction, testAction)
		})
	}
}
