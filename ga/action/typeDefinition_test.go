package action_test

import (
	"go-actions/ga/action"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
	"reflect"
	"testing"
)

var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockGenerator}
var mockActionConfig = &action.ActionConfig{}

func TestTypeDefinitionFromRegistration(t *testing.T) {
	reg := action.ActionRegistration[ta.ActionValidEmpty]{Action: ta.ActionValidEmpty{}}

	expectedTypeName := "ActionValidEmpty"
	expectedTypePath := "go-actions/ga/utils/testing/testActions/testActions.ActionValidEmpty"
	expectedType := reflect.TypeOf(ta.ActionValidEmpty{})
	expectedValue := reflect.ValueOf(ta.ActionValidEmpty{})
	// expectedPropsType := reflect.TypeOf(reg.DefaultProps)
	// expectedPropsValue := reflect.ValueOf(reg.DefaultProps)
	expectedTriggerValue := false

	defReg := action.TypeDefinitionFromRegistration(&reg)

	assert.Equals(t, expectedTypeName, defReg.TypeName)
	assert.Equals(t, expectedTypePath, defReg.TypePath)
	assert.Equals(t, expectedType, defReg.ActionType)
	assert.Equals(t, expectedValue, defReg.ActionValue)
	// assert.Equals(t, expectedPropsType, defReg.PropsType)
	// assert.Equals(t, expectedPropsValue, defReg.PropsValue)
	assert.Equals(t, expectedTriggerValue, defReg.Trigger)
}

func TestTriggerDefinitionFromRegistration(t *testing.T) {
	reg := action.ActionRegistration[ta.ActionTriggerValid]{Action: ta.ActionTriggerValid{}}
	defReg := action.TypeDefinitionFromRegistration(&reg)

	assert.Equals(t, true, defReg.Trigger)
}

// func TestNewDefaultProps(t *testing.T) {
// 	reg := ta.GenerateActionValidRegistration()
// 	defReg := TypeDefinitionFromRegistration(&reg)

// 	newProps := defReg.NewDefaultProps()
// 	typeAssertedProps, ok := newProps.(ta.ActionValidProps)

// 	assert.Equals(t, true, ok)
// 	assert.Equals(t, reg.DefaultProps, typeAssertedProps)
// }

// func TestValidatePropsType(t *testing.T) {
// 	reg := ta.GenerateActionValidRegistration()
// 	defReg := TypeDefinitionFromRegistration(&reg)

// 	tests := []struct {
// 		name      string
// 		props     action.GoActionProps
// 		expectErr bool
// 	}{
// 		{name: "valid - same props type", props: ta.ActionValidProps{}, expectErr: false},
// 		{name: "valid - same props type generated", props: defReg.NewDefaultProps(), expectErr: false},
// 		{name: "valid - same props type populated", props: ta.ActionValidProps{Param1: "asdf"}, expectErr: false},
// 		{name: "invalid - same props pointer", props: &ta.ActionValidProps{}, expectErr: true},
// 		{name: "invalid - different props type", props: ta.ActionInvalidNoExecute{}, expectErr: true},
// 		{name: "invalid - nil", props: nil, expectErr: true},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			err := defReg.ValidatePropsType(test.props)

// 			assert.Equals(t, test.expectErr, err != nil)
// 		})
// 	}
// }

func TestNewAction(t *testing.T) {
	reg := action.ActionRegistration[ta.ActionValid]{Action: ta.ActionValid{}}
	defReg := action.TypeDefinitionFromRegistration(&reg)

	inst := action.NewActionInstance("", mockGlobalConfig, mockActionConfig)

	act, err := action.NewAction[ta.ActionValid](defReg, inst)

	assert.Equals(t, nil, err)
	assert.Equals(t, ta.ActionValid{}, act)
}
