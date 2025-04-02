package flow

import (
	"go-actions/ga/app"
	"go-actions/ga/testing/assert"
	"go-actions/ga/testing/testActions"
	"testing"
)

func TestInitFlow(t *testing.T) {
	app := app.NewApp("test")
	def := NewFlowDefinition()
	flow := NewFlow(app, def)

	if flow == nil {
		t.Errorf("expected type of %v but got %v", Flow{}, nil)
	}
}

func TestAddAction(t *testing.T) {
	type test struct {
		name             string
		actionRegistered bool
		expectedActions  int
		err              bool
	}

	cases := []test{
		{name: "registered action", actionRegistered: true, expectedActions: 1, err: false},
		{name: "unregistered action", actionRegistered: false, expectedActions: 0, err: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reg := testActions.GenerateActionValidEmptyRegistration()
			a := app.NewApp("test")

			if tc.actionRegistered {
				app.RegisterAction(&reg)(a)
			}
			def := NewFlowDefinition()
			f := NewFlow(a, def)
			_, err := NewFlowAction[testActions.ActionValidEmpty](f, &testActions.ActionValidEmptyProps{})
			assert.Equals(t, tc.expectedActions, len(f.flowDefinition.Actions))
			assert.Equals(t, tc.err, err != nil)
		})
	}
}
