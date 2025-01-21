package flow

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/app"
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testActions"
	"testing"
)

func TestInitFlow(t *testing.T) {
	app := app.NewApp("test")
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
			reg := testActions.GenerateActionValidEmptyRegistration()
			a := app.NewApp("test")

			if tc.actionRegistered {
				app.RegisterAction(&reg)(a)
			}

			defer func() {
				didPanic := recover() != nil
				asserts.Equals(t, tc.expectPanic, didPanic)
			}()

			f := NewFlow(a)
			AddAction[testActions.ActionValidEmpty](&testActions.ActionValidEmptyProps{})(f)
			asserts.Equals(t, tc.expectedActions, len(f.actionInstances))
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	flowApp := app.NewApp("test")
	reg := testActions.GenerateActionValidEmptyRegistration()
	app.RegisterAction(&reg)(flowApp)

	flow := NewFlow(flowApp)
	action := AddAction[testActions.ActionValidEmpty](&testActions.ActionValidEmptyProps{})(flow)

	marshalledActionInstance, _ := json.Marshal(action.Instance)
	expected := fmt.Sprintf(`{"actions":[%s]}`, string(marshalledActionInstance))

	marshalledFlow, _ := json.Marshal(flow)

	asserts.Equals(t, string(expected), string(marshalledFlow))

}
