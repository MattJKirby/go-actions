package model

import (
	"encoding/json"
	"go-actions/ga/utils/testing/assert"
	"go-actions/ga/utils/testing/testHelpers/actionTestHelpers"
	"testing"
)

var mockConfig = &actionTestHelpers.MockActionConfig{MockUid: "uid"}

func TestMarshalEmptyModel(t *testing.T) {
	model := NewActionModel("someName", mockConfig)
	mashalled, _ := json.Marshal(model)

	assert.Equals(t, `{"name":"someName","uid":"someName:uid","parameters":[],"inputs":[],"outputs":[]}`, string(mashalled))
}

func TestUnmarshalmodel(t *testing.T) {
	model := NewActionModel("someName", mockConfig)
	marshalled := `{"name":"otherName","uid":"otherUid","parameters":[],"inputs":[],"outputs":[]}`

	err := json.Unmarshal([]byte(marshalled), model)
	assert.Equals(t, err, nil)
	assert.Equals(t, model.ActionName, "otherName")
	assert.Equals(t, model.ActionUid, "otherUid")
}
