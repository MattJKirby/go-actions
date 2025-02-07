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

	defReg := TypeDefinitionFromRegistration(&reg)

	asserts.Equals(t, expectedTypeName, defReg.TypeName)
	asserts.Equals(t, expectedTypePath, defReg.TypePath)
	asserts.Equals(t, expectedType, defReg.ActionType)
	asserts.Equals(t, expectedValue, defReg.ActionValue)
	asserts.Equals(t, expectedCtor, defReg.CtorValue.Pointer())
	asserts.Equals(t, expectedCtorType, defReg.CtorType)
	asserts.Equals(t, expectedPropsType, defReg.PropsType)
	asserts.Equals(t, expectedPropsValue, defReg.PropsValue)
}

func TestNewDefaultProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	newProps := defReg.NewDefaultProps()
	typeAssertedProps, ok := newProps.(ta.ActionValidProps)

	asserts.Equals(t, true, ok)
	asserts.Equals(t, *reg.DefaultProps, typeAssertedProps)
}

func TestNewConstructorWithValidProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)

	expectedInst := action.NewActionInstance("expected", mockConfig)
	expectedAction := reg.Constructor(expectedInst, ta.ActionValidDefaultProps)

	testInst := action.NewActionInstance("test", mockConfig)
	testCtor := defReg.NewConstructor()

	abc := defReg.NewDefaultProps()
	testAction, err := testCtor(testInst, abc)
	typedTestAction, ok := (testAction).(*ta.ActionValid)

	asserts.Equals(t, nil, err)
	asserts.Equals(t, true, ok)
	asserts.Equals(t, expectedAction, typedTestAction)
	asserts.Equals(t, expectedInst.Model.Parameters, testInst.Model.Parameters)

}

func TestNewConstructorInvalidProps(t *testing.T) {
	reg := ta.GenerateActionValidRegistration()
	defReg := TypeDefinitionFromRegistration(&reg)
	emptyInst := action.NewActionInstance("expected", mockConfig)

	tests := []struct {
		name      string
		input     any
		expectErr bool
	}{
		{name: "construct with matching value prop type", input: &ta.ActionValidProps{}, expectErr: true},
		{name: "construct with wrong prop type", input: ta.ActionValidEmptyProps{}, expectErr: true},
		{name: "construct with nil prop type", input: nil, expectErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testInst := action.NewActionInstance("test", mockConfig)
			testCtor := defReg.NewConstructor()
			testAction, err := testCtor(testInst, test.input)

			asserts.Equals(t, test.expectErr, err != nil)
			asserts.Equals(t, nil, testAction)
			asserts.Equals(t, emptyInst.Model.Parameters, testInst.Model.Parameters)
		})
	}

}
