package flow

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/app"
	"go-actions/ga/app/config"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testActions"
	"go-actions/ga/utils/testing/testHelpers"
	"testing"
)

var mockGenerator = &testHelpers.MockUidGenerator{MockUid: "uid"}
var mockConfig = &config.GlobalConfig{UidGenerator: mockGenerator}

func TestAddInstance(t *testing.T) {
	testcases := []struct{
		name string
		alreadyExists bool
		err bool
	}{
		{name: "valid - instance does not already exist", alreadyExists: false, err: false},
    {name: "invalid - instance already exists", alreadyExists: true, err: true},
	}

  for _, test := range testcases {
    t.Run(test.name, func(t *testing.T) {
      a := app.NewApp("testApp")
			flowDef := NewFlowDefinition(a)
      instance := action.NewActionInstance("someInstance", mockConfig)

      if test.alreadyExists {
        flowDef.addInstance(instance)
      }

      err := flowDef.addInstance(instance)
      assert.Equals(t, test.err, err != nil)
    })
  }
}

func TestNewAction(t *testing.T) {
	a := app.NewApp("testApp")
	reg := testActions.GenerateActionValidEmptyRegistration()
	app.RegisterAction(&reg)(a)

	flowDef := NewFlowDefinition(a)
	act, err := flowDef.NewAction("ActionValidEmpty")

	assert.Equals(t, true, flowDef.Actions[act.Instance.Model.ActionUid] != nil)
	assert.Equals(t, nil, err)
}

func TestMarshalFlowDefinition(t *testing.T) {
	a := app.NewApp("testApp")
	flowDef := NewFlowDefinition(a)
	instance := action.NewActionInstance("someInstance", mockConfig)

	flowDef.addInstance(instance)
	marshalledInstance, _ := json.Marshal(instance)

	marshalled, err := json.Marshal(flowDef)
	expected := fmt.Sprintf(`{"Actions":{"%s":%s}}`, instance.Model.ActionUid, marshalledInstance)
	assert.Equals(t, nil, err)
	assert.Equals(t, expected, string(marshalled))
}
