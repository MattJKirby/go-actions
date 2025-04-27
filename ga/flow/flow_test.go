package flow

import (
	"go-actions/ga/app"
	"testing"
)

func TestInitFlow(t *testing.T) {
	app := app.NewApp("test")
	flow := NewFlow(app)

	if flow == nil {
		t.Errorf("expected type of %v but got %v", Flow{}, nil)
	}
}
