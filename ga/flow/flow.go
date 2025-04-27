package flow

import (
	"go-actions/ga/app"
)

type Flow struct {
	flowApp        *app.App
	flowDefinition *flowDefinition
}

func NewFlow(app *app.App) *Flow {
	return &Flow{
		flowApp:        app,
		flowDefinition: NewFlowDefinition(),
	}
}
