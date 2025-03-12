package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/model"
	"go-actions/ga/utils"
	"reflect"
)

type App struct {
	Name           string
	ctx            context.Context
	actionRegistry *actionRegistry
	modelConfig    model.ActionModelConfig
}

func NewApp(name string) *App {
	return &App{
		Name:           name,
		ctx:            context.Background(),
		actionRegistry: newActionRegistry(),
		modelConfig:    model.NewModelConfig(),
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

func GetActionByName(actionName string) func(*App) (*InitialisedAction, error) {
	return func(app *App) (*InitialisedAction, error) {
		typeDef, err := getRegisteredTypeDefinitionByName(actionName)(app.actionRegistry)
		if err != nil {
			return nil, err
		}
		return InitialiseNewAction(app.modelConfig, typeDef)
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*InitialisedTypedAction[T], error) {
	return func(app *App) (*InitialisedTypedAction[T], error) {
		def, err := GetDefinitionByType[T, P]()(app)
		if err != nil {
			return nil, err
		}

		return InitialiseNewTypedAction(app.modelConfig, def)
	}
}
