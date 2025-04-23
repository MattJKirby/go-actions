package flow

import (
	"go-actions/ga/app"
)

type Flow struct {
	flowApp        *app.App
	flowDefinition *flowDefinition
}

func NewFlow(app *app.App, flowDefinition *flowDefinition) *Flow {
	return &Flow{
		flowApp:        app,
		flowDefinition: flowDefinition,
	}
}
