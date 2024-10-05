package app

import (
	"go-actions/ga/action"
	"testing"
)

func TestRegisterDef(t *testing.T){
	app := NewApp()
	def := action.ActionDefinition{}

	result := app.RegisterActionDef(&def)

	t.Run("Test register def", func(t *testing.T) {
		if result != &def {
			t.Errorf("Error during registration: expected %v, got %v", &def, result)
		}
	})
}