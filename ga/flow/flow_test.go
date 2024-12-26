package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/cr/asserts"
	"testing"
)

type testAction struct{}

func (ta testAction) Execute() {}
func testActionCtor(*action.ActionInstance) *testAction {
	return &testAction{}
}

func TestInitFlow(t *testing.T) {
	app := app.NewApp()
	flow := NewFlow(app)

	if flow == nil {
		t.Errorf("expected type of %v but got %v", Flow{}, nil)
	}
}

func TestAddAction(t *testing.T) {
	type test struct {
		name             string
		actionRegistered bool
		expectedActions  int
		expectPanic      bool
	}

	cases := []test{
		{name: "registered action", actionRegistered: true, expectedActions: 1, expectPanic: false},
		{name: "unregistered action", actionRegistered: false, expectedActions: 0, expectPanic: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := app.NewApp()

			if tc.actionRegistered {
				app.RegisterAction(testActionCtor)(a)
			}

			defer func() {
				didPanic := recover() != nil
				asserts.Equals(t, tc.expectPanic, didPanic)
			}()

			f := NewFlow(a)
			AddAction(testAction{})(f)
			asserts.Equals(t, tc.expectedActions, len(f.actions))
		})
	}

}
