package model

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

type cnfg struct{}

func (c *cnfg) GenerateUid() string {
	return "uid"
}

func TestNewModelInstance(t *testing.T) {
	instance := NewModelInstance("someName", &cnfg{})
	
	asserts.Equals(t, "someName:uid", instance.ActionUid)
}

func TestMarshalEmptyModel(t *testing.T) {
	instance := NewModelInstance("someName", &cnfg{})
	mashalled, _ := json.Marshal(instance)

	asserts.Equals(t, `{"name":"someName","uid":"someName:uid","parameters":{},"inputs":{}}`, string(mashalled))
}

func TestUnmarshalInstance(t *testing.T) {
	instance := NewModelInstance("someName", &cnfg{})
	marshalled := `{"name":"otherName","uid":"otherUid","parameters":{},"inputs":{}}`

	err := json.Unmarshal([]byte(marshalled), instance)
	asserts.Equals(t, err, nil)
	asserts.Equals(t, instance.ActionName, "otherName")
	asserts.Equals(t, instance.ActionUid, "otherUid")
}
