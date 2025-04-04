package app

import (
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/utils/testing/assert"
	ta "go-actions/ga/utils/testing/testActions"
	"testing"
)

func appWithValidActionRegistration() (*App, action.GoActionRegistration[ta.ActionValid, ta.ActionValidProps]) {
	app := NewApp("test")
	app.actionConfig = mockConfig
	registration := ta.GenerateActionValidRegistration()
	RegisterAction(&registration)(app)
	return app, registration
}

func TestInitialiseNewAction(t *testing.T) {
	reg := ta.GenerateActionValidEmptyRegistration()
	definition := definition.TypeDefinitionFromRegistration(&reg)
	instance := action.NewActionInstance("ActionValidEmpty", mockConfig)

	expected := &InitialisedAction{
		Action:              reg.Constructor(instance, ta.ActionValidEmptyProps{}),
		InitialisedInstance: instance,
	}

	actual, err := InitialiseNewAction(mockConfig, definition)
	assert.Equals(t, expected, actual)
	assert.Equals(t, nil, err)
}

func TestInitialiseTypedAction(t *testing.T) {
	app, reg := appWithValidActionRegistration()
	def := definition.NewActionDefinition(&reg)

	tests := []struct {
		name  string
		props *ta.ActionValidProps
	}{
		{name: "init with default props", props: reg.DefaultProps},
		{name: "init with non default props", props: &ta.ActionValidProps{Param1: "Some Val"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			instance := action.NewActionInstance(def.TypeName, mockConfig)
			expectedInstantiatedTypedAction := &InitialisedTypedAction[ta.ActionValid]{
				Action:              reg.Constructor(instance, *test.props),
				InitialisedInstance: instance,
			}

			actual, err := InitialiseNewTypedAction(app.actionConfig, def, test.props)

			assert.Equals(t, expectedInstantiatedTypedAction, actual)
			assert.Equals(t, nil, err)
		})
	}
}
