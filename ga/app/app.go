package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"

	"go-actions/ga/app/registration"
	"go-actions/ga/libs/packageConfig"

	"reflect"
)

type App struct {
	Name           string
	ctx            context.Context
	Config         *ApplicationConfig
	actionRegistry *registration.ActionRegistry
}

func NewApp(name string, opts ...packageConfig.Option[ApplicationConfig]) *App {
	return &App{
		Name:           name,
		ctx:            context.Background(),
		Config:         packageConfig.NewPackageConfig(DefaultApplicationConfig(), opts...),
		actionRegistry: registration.NewActionRegistry(),
	}
}

func RegisterAction[T action.GoAction](reg *action.ActionRegistration[T]) func(*App) {
	return func(app *App) {
		if reg != nil {
			registration.AcceptRegistration(reg)(app.actionRegistry)
			return
		}
		registration.AcceptRegistration(&action.ActionRegistration[T]{Action: *new(T)})(app.actionRegistry)
	}
}

func GetDefinitionByType[T action.GoAction]() func(*App) (*action.TypeDefinition, error) {
	return func(app *App) (*action.TypeDefinition, error) {
		action := new(T)
		actionType := reflect.TypeOf(*action)
		return registration.GetTypeDefinitionByType(actionType)(app.actionRegistry)
	}
}

func GetDefinitionByName(name string) func(*App) (*action.TypeDefinition, error) {
	return func(app *App) (*action.TypeDefinition, error) {
		return registration.GetTypeDefinitionByName(name)(app.actionRegistry)
	}
}

func GetActionByName(typeDef *action.TypeDefinition, inst *action.ActionInstance) func(*App) (*executable.Action[action.GoAction], error) {
	return func(app *App) (*executable.Action[action.GoAction], error) {
		if inst == nil {
			inst = action.NewActionInstance(app.Config.Global, app.Config.Action, typeDef)
		}

		return executable.NewAction[action.GoAction](typeDef, inst)
	}
}

