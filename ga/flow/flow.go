package flow

import (
	"go-actions/ga/app"
)

type Flow struct {
	flowApp    *app.App
	Definition *flowDefinition
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:    app,
		Definition: NewFlowDefinition(app),
	}
}
