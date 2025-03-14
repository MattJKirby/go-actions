package io

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewOutput(t *testing.T) {
	output := NewActionOutput("name", "uid")
	asserts.Equals(t, "name", output.Name)
	asserts.Equals(t, "uid__Output:name", output.Id)
	asserts.Equals(t, 0, len(output.TargetReferences))
}

func TestAssignInputRef(t *testing.T) {
	output := NewActionOutput("name", "uid")
	inputRef := NewReference("other uid", "output")

	output.AssignTarget(inputRef)
	asserts.Equals(t, 1, len(output.TargetReferences))
	asserts.Equals(t, inputRef, output.TargetReferences[0])
}

func TestMarshal(t *testing.T) {
	output := NewActionOutput("o", "uid")
	t.Run("no inputs", func(t *testing.T) {
		marshalled, err := json.Marshal(output)
		asserts.Equals(t, nil, err)
		asserts.Equals(t, `{"name":"o","id":"uid__Output:o","targets":[]}`, string(marshalled))
	})

	t.Run("with inputs", func(t *testing.T) {
		inputRef := NewReference("other uid", "output")
		output.AssignTarget(inputRef)
		marshalledInputRef, _ := json.Marshal(inputRef)
		marshalledOutput, err := json.Marshal(output)

		expected := fmt.Sprintf(`{"name":"o","id":"uid__Output:o","targets":[%s]}`, string(marshalledInputRef))
		asserts.Equals(t, nil, err)
		asserts.Equals(t, expected, string(marshalledOutput))
	})
}

func TestUnmarshal(t *testing.T) {
	output := NewActionOutput("o", "uid")
	inputRef := NewReference("uid", "input")
	output.AssignTarget(inputRef)

	t.Run("valid with inputs", func(t *testing.T) {
		marshalled, _ := json.Marshal(output)
		newOutput := NewActionOutput("o", "uid")
		json.Unmarshal(marshalled, newOutput)

		asserts.Equals(t, output.Name, newOutput.Name)
		asserts.Equals(t, output.Id, newOutput.Id)
		asserts.Equals(t, len(output.TargetReferences), len(newOutput.TargetReferences))
	})

	t.Run("invalid", func(t *testing.T) {
		marshalled, _ := json.Marshal(output)
		newOutput := NewActionOutput("badName", "uid")
		err := json.Unmarshal(marshalled, newOutput)
		if err == nil {
			t.Errorf("expected err but got %v", nil)
		}
	})
}
