package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestInitFlow(t *testing.T) {
	flow := NewFlow()

	if flow == nil {
		t.Errorf("expected type of %v but got %v", Flow{}, nil)
	}

	if flow.actions == nil {
		t.Errorf("error initialising flow actions: expected map but got %v", nil)
	}
}

type testAction struct{}

func (ta testAction) Execute() {}
func testActionCtor(*action.ActionInstance) *testAction {
	return &testAction{}
}

func TestNewAction(t *testing.T) {
	a := app.NewApp()
	app.RegisterAction(testActionCtor)(a)
	flow := NewFlow()
	action, _ := app.NewAction[testAction](testAction{})(a)
	result := NewAction(action)(flow)

	asserts.Equals(t, 1, len(flow.actions))
	asserts.Equals(t, action, result)
}
