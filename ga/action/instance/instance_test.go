package instance

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

type cnfg struct{}

func (c *cnfg) GenerateUid() string {
	return "uid"
}

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("someName", &cnfg{})
	t.Run("test new instance", func(t *testing.T) {
		asserts.Equals(t, "someName:uid", instance.ActionUid)
	})
}

func TestMarshalEmptyInstance(t *testing.T) {
	instance := NewActionInstance("someName", &cnfg{})
	mashalled, _ := json.Marshal(instance)
	t.Run("empty instance", func(t *testing.T) {
		asserts.Equals(t, `{"name":"someName","uid":"someName:uid","parameters":{},"inputs":{}}`, string(mashalled))
	})
}

func TestUnmarshalInstance(t *testing.T) {
	instance := NewActionInstance("someName", &cnfg{})
	marshalled := `{"name":"otherName","uid":"otherUid","parameters":{},"inputs":{}}`

	t.Run("test unmarsh", func(t *testing.T) {
		err := json.Unmarshal([]byte(marshalled), instance)
		asserts.Equals(t, err, nil)
		asserts.Equals(t, instance.ActionName, "otherName")
		asserts.Equals(t, instance.ActionUid, "otherUid")
	})
}
