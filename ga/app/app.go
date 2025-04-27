package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app/config"
	"go-actions/ga/app/registration"
	"go-actions/ga/utils/packageConfig"

	"reflect"
)

type App struct {
	Name           string
	ctx            context.Context
	config         *config.ApplicationConfig
	actionRegistry *registration.ActionRegistry
}

func NewApp(name string, opts ...packageConfig.Option[config.ApplicationConfig]) *App {
	return &App{
		Name:           name,
		ctx:            context.Background(),
		config:         packageConfig.NewPackageConfig(config.DefaultApplicationConfig(), opts...),
		actionRegistry: registration.NewActionRegistry(),
	}
}

func RegisterAction[T action.GoAction](reg *action.GoActionRegistration[T]) func(*App) {	
	return func(app *App) {
		if reg != nil {
			registration.AcceptRegistration(reg)(app.actionRegistry)
			return
		}
		registration.AcceptRegistration(&action.GoActionRegistration[T]{Action: *new(T)})
	}
}

func GetDefinitionByType[T action.GoAction]() func(*App) (*definition.ActionTypeDefinition, error) {
	return func(app *App) (*definition.ActionTypeDefinition, error) {
		action := new(T)
		actionType := reflect.TypeOf(*action)
		return registration.GetTypeDefinitionByType(actionType)(app.actionRegistry)
	}
}

func GetDefinitionByName(name string) func(*App) (*definition.ActionTypeDefinition, error) {
	return func(app *App) (*definition.ActionTypeDefinition, error) {
		return registration.GetTypeDefinitionByName(name)(app.actionRegistry)
	}
}

func GetActionByName(actionName string) func(*App) (*executable.Action[action.GoAction], error) {
	return func(app *App) (*executable.Action[action.GoAction], error) {
		typeDef, err := registration.GetTypeDefinitionByName(actionName)(app.actionRegistry)
		if err != nil {
			return nil, err
		}
		return executable.NewAction[action.GoAction](app.config.Global, typeDef)
	}
}

func GetAction[T action.GoAction]() func(*App) (*executable.Action[T], error) {
	return func(app *App) (*executable.Action[T], error) {
		def, err := GetDefinitionByType[T]()(app)
		if err != nil {
			return nil, err
		}
		return executable.NewAction[T](app.config.Global, def)
	}
}
