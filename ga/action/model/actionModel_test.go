package model

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

type mockConfig struct {
	uid string
}

func (mc mockConfig) GenerateUid() string {
	return mc.uid
}

func TestNewModelInstance(t *testing.T) {
	model := NewActionModel("someName", &mockConfig{"uid"})

	asserts.Equals(t, "someName:uid", model.ActionUid)
}

func TestMarshalEmptyModel(t *testing.T) {
	model := NewActionModel("someName", &mockConfig{"uid"})
	mashalled, _ := json.Marshal(model)

	asserts.Equals(t, `{"name":"someName","uid":"someName:uid","parameters":[],"inputs":[],"outputs":[]}`, string(mashalled))
}

func TestUnmarshalmodel(t *testing.T) {
	model := NewActionModel("someName", &mockConfig{"uid"})
	marshalled := `{"name":"otherName","uid":"otherUid","parameters":[],"inputs":[],"outputs":[]}`

	err := json.Unmarshal([]byte(marshalled), model)
	asserts.Equals(t, err, nil)
	asserts.Equals(t, model.ActionName, "otherName")
	asserts.Equals(t, model.ActionUid, "otherUid")
}
