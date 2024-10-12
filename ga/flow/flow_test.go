package flow

import (
	"testing"
)

func TestInitFlow(t *testing.T) {
	flow := NewFlow()

	t.Run("Assert new flow and properties", func(t *testing.T){
		if flow == nil {
			t.Errorf("expected type of %v but got %v", Flow{}, nil)
		}

		if flow.actions == nil {
			t.Errorf("error initialising flow actions: expected map but got %v", nil)
		}

	})
}