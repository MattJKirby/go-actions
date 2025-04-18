package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
	"go-actions/ga/app/config"
	"go-actions/ga/utils"
	"go-actions/ga/utils/packageConfig"

	"reflect"
)

type App struct {
	Name           string
	ctx            context.Context
	config         *config.ApplicationConfig
	actionRegistry *actionRegistry
}

func NewApp(name string, opts ...packageConfig.Option[config.ApplicationConfig]) *App {
	return &App{
		Name:           name,
		ctx:            context.Background(),
		config:         packageConfig.NewPackageConfig(config.DefaultApplicationConfig(), opts...),
		actionRegistry: newActionRegistry(),
	}
}

func RegisterAction[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) func(*App) {
	return func(app *App) {
		acceptRegistration(reg)(app.actionRegistry)
	}
}

func GetDefinitionByType[T action.GoAction, P action.GoActionProps]() func(*App) (*definition.ActionDefinition[T, P], error) {
	return func(app *App) (*definition.ActionDefinition[T, P], error) {
		action := new(T)
		actionType := utils.GetValueType(reflect.TypeOf(action))
		return getTypedActionDefinition[T, P](actionType)(app.actionRegistry)
	}
}

func GetDefinitionByName(name string) func(*App) (*definition.ActionTypeDefinition, error) {
	return func(app *App) (*definition.ActionTypeDefinition, error) {
		return getRegisteredTypeDefinitionByName(name)(app.actionRegistry)
	}
}

func GetActionByName(actionName string) func(*App) (*executable.BaseExecutable[action.GoAction], error) {
	return func(app *App) (*executable.BaseExecutable[action.GoAction], error) {
		typeDef, err := getRegisteredTypeDefinitionByName(actionName)(app.actionRegistry)
		if err != nil {
			return nil, err
		}
		return executable.NewBaseExecutable[action.GoAction](app.config.Global, typeDef)
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*executable.BaseExecutable[T], error) {
	return func(app *App) (*executable.BaseExecutable[T], error) {
		def, err := GetDefinitionByType[T, P]()(app)
		if err != nil {
			return nil, err
		}
		return executable.NewBaseExecutable[T](app.config.Global, def.GetTypeDefinition())
	}
}
