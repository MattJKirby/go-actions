package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/action/model/io/reference"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewOutput(t *testing.T) {
	output := NewActionOutput("name", "uid")
	asserts.Equals(t, "name", output.Name)
	asserts.Equals(t, "uid__Output:name", output.Id)
	asserts.Equals(t, 0, len(output.InputReferences))
}

func TestAssignInputRef(t *testing.T) {
	output := NewActionOutput("name", "uid")
	inputRef := reference.NewInputReference("other uid", "output")

	output.AssignInputReference(inputRef)
	asserts.Equals(t, 1, len(output.InputReferences))
	asserts.Equals(t, inputRef, output.InputReferences[0])
}

func TestMarshal(t *testing.T) {
	output := NewActionOutput("o", "uid")
	t.Run("no inputs", func(t *testing.T) {
		marshalled, err := json.Marshal(output)
		asserts.Equals(t, nil, err)
		asserts.Equals(t, `{"name":"o","id":"uid__Output:o","inputs":[]}`, string(marshalled))
	})

	t.Run("with inputs", func(t *testing.T) {
		inputRef := reference.NewInputReference("other uid", "output")
		output.AssignInputReference(inputRef)
		marshalledInputRef, _ := json.Marshal(inputRef)
		marshalledOutput, err := json.Marshal(output)

		expected := fmt.Sprintf(`{"name":"o","id":"uid__Output:o","inputs":[%s]}`, string(marshalledInputRef))
		asserts.Equals(t, nil, err)
		asserts.Equals(t, expected, string(marshalledOutput))
	})
}
