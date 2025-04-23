package flow

import (
	"go-actions/ga/app"
	"testing"
)

func TestInitFlow(t *testing.T) {
	app := app.NewApp("test")
	def := NewFlowDefinition()
	flow := NewFlow(app, def)

	if flow == nil {
		t.Errorf("expected type of %v but got %v", Flow{}, nil)
	}
}
