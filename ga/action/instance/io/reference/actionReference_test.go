package reference

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestMarshalReference(t *testing.T) {
	outputRef := NewReference("a", "output", "type")

	t.Run("marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(outputRef)
		asserts.Equals(t, `{"actionUid":"a","resourceName":"output"}`, string(marshalled))
	})
}

func TestUnmarshalReference(t *testing.T) {
	marshalled := []byte(`{"actionUid":"a","resourceName":"output"}`)
	ref := NewReference("", "", "")

	t.Run("unmarshal", func(t *testing.T) {
		json.Unmarshal(marshalled, ref)

		asserts.Equals(t, "a", ref.ActionUid)
		asserts.Equals(t, "output", ref.ResourceName)
	})
}
