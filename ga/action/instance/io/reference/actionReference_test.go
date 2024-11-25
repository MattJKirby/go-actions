package reference

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestMarshalReference(t *testing.T) {
	outputRef := NewOutputReference("a", "o")

	t.Run("marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(outputRef)
		asserts.Equals(t, `{"actionUid":"a","resourceName":"o"}`, string(marshalled))
	})
}

func TestUnmarshalReference(t *testing.T) {
	marshalled := []byte(`{"actionUid":"a","resourceName":"o"}`)
	ref := NewOutputReference("", "")

	t.Run("unmarshal", func(t *testing.T) {
		json.Unmarshal(marshalled, ref)

		asserts.Equals(t, "a", ref.ActionUid)
		asserts.Equals(t, "o", ref.ResourceName)
	})
}
