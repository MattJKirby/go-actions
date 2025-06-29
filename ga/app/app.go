package app

import (
	"context"
	"go-actions/ga/action"

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

func RegisterAction[T action.GoAction](act T, reg *action.ActionRegistration[T]) func(*App) {
	return func(app *App) {
		if reg != nil {
			registration.AcceptRegistration(reg)(app.actionRegistry)
			return
		}
		registration.AcceptRegistration(&action.ActionRegistration[T]{Action: act})(app.actionRegistry)
	}
}

func GetDefinitionByType[T action.GoAction](a T) func(*App) (*action.TypeDefinition, error) {
	return func(app *App) (*action.TypeDefinition, error) {
		actionType := reflect.TypeOf(a)
		return registration.GetTypeDefinitionByType(actionType)(app.actionRegistry)
	}
}

func GetDefinitionByName(name string) func(*App) (*action.TypeDefinition, error) {
	return func(app *App) (*action.TypeDefinition, error) {
		return registration.GetTypeDefinitionByName(name)(app.actionRegistry)
	}
}

func GetAction[T action.GoAction](typeDef *action.TypeDefinition, inst *action.ActionInstance) func(*App) (*action.Action[T], error) {
	return func(app *App) (*action.Action[T], error) {
		if inst == nil {
			inst = action.NewActionInstance(app.Config.Global, typeDef)
		}

		return action.NewAction[T](typeDef, inst)
	}
}
