package reference

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewActionReference(t *testing.T) {
	ref := NewReference("actionId", "recieverId", "refType")

	t.Run("assert reference id", func(t *testing.T) {
		asserts.Equals(t, "actionId__ref:refType:recieverId", ref.id)
	})
}

func TestMarshalReference(t *testing.T) {
	outputRef := NewReference("a", "output", "type")

	t.Run("marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(outputRef)
		asserts.Equals(t, `{"actionUid":"a","resourceName":"output"}`, string(marshalled))
	})
}
