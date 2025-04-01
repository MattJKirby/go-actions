package model

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"go-actions/ga/testing/testHelpers/actionTestHelpers"
	"testing"
)

var mockConfig = &actionTestHelpers.MockActionConfig{MockUid: "uid"}

func TestMarshalEmptyModel(t *testing.T) {
	model := NewActionModel("someName", mockConfig)
	mashalled, _ := json.Marshal(model)

	asserts.Equals(t, `{"name":"someName","uid":"someName:uid","parameters":[],"inputs":[],"outputs":[]}`, string(mashalled))
}

func TestUnmarshalmodel(t *testing.T) {
	model := NewActionModel("someName", mockConfig)
	marshalled := `{"name":"otherName","uid":"otherUid","parameters":[],"inputs":[],"outputs":[]}`

	err := json.Unmarshal([]byte(marshalled), model)
	asserts.Equals(t, err, nil)
	asserts.Equals(t, model.ActionName, "otherName")
	asserts.Equals(t, model.ActionUid, "otherUid")
}
