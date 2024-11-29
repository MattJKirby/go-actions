package instance

import (
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

