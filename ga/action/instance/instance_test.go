package instance

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

type cnfg struct {}
func (c *cnfg) GenerateUid() string {
	return "uid"
}

func TestNewActionInstance(t *testing.T) {
	instance := NewActionInstance("someName", &cnfg{})
	t.Run("test new instance", func(t *testing.T) {
		asserts.Equals(t, "someName:uid", instance.ActionUid)
	})
}

func TestMarshalEmptyInstance(t *testing.T){
	instance := NewActionInstance("someName", &cnfg{})
	mashalled, _ := json.Marshal(instance)
	t.Run("empty instance", func(t *testing.T) {
		asserts.Equals(t, `{"name":"someName","uid":"someName:uid","parameters":{},"inputs":{}}`, string(mashalled))
	})
}
