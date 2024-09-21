package app

import "context"

type App struct {
	ctx context.Context
	actionDefinitionRegistry *ActionDefinitionRegistry
}

func NewApp() *App{
	return &App{
		ctx: context.Background(),
		actionDefinitionRegistry: NewActionDefinitionRegistry(),
	}
}