package reference

import (
	"encoding/json"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestMarshalOutputReference(t *testing.T) {
	outputRef := NewOutputReference("a", "o")
	marshalledOutput, _ := json.Marshal(outputRef)
	asserts.Equals(t, `{"actionUid":"a","resourceName":"o"}`, string(marshalledOutput))
}

func TestMarshalInputReference(t *testing.T) {
	inputRef := NewInputReference("a", "i")
	marshalledInput, _ := json.Marshal(inputRef)
	asserts.Equals(t, `{"actionUid":"a","resourceName":"i"}`, string(marshalledInput))
}

func TestUnmarshalOutputReference(t *testing.T) {
	marshalledOutput := []byte(`{"actionUid":"a","resourceName":"o"}`)
	outputRef := NewOutputReference("", "")

	json.Unmarshal(marshalledOutput, outputRef)
	asserts.Equals(t, "a", outputRef.ActionUid)
	asserts.Equals(t, "o", outputRef.ResourceName)
}

func TestUnmarshalInputReference(t *testing.T) {
	marshalledInput := []byte(`{"actionUid":"a","resourceName":"i"}`)
	inputRef := NewInputReference("", "")

	json.Unmarshal(marshalledInput, inputRef)
	asserts.Equals(t, "a", inputRef.ActionUid)
	asserts.Equals(t, "i", inputRef.ResourceName)
}
