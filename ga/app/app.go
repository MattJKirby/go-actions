package app

import (
	"context"
	"fmt"
	"go-actions/ga/action"
	"go-actions/ga/action/definition"
	"go-actions/ga/action/executable"
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

func GetActionRegistration[T action.GoAction, P action.GoActionProps]() func(*App) (*definition.ActionDefinition[T, P], error) {
	return func(app *App) (*definition.ActionDefinition[T, P], error) {
		action := new(T)
		actionType := utils.GetValueType(reflect.TypeOf(action))
		return getTypedActionDefinition[T, P](actionType)(app.actionRegistry)
	}
}

func InstantiateAction(actionName string) func(*App) (*executable.ExecutableAction, error) {
	return func(app *App) (*executable.ExecutableAction, error) {
		typeDef, err := getRegisteredTypeDefinitionByName(actionName)(app.actionRegistry)
		if err != nil {
			return nil, err
		}
		return executable.NewExecutableAction(app.modelConfig, typeDef), nil
	}
}

func InstantiateTypedAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*executable.TypedExecutable[T, P], error) {
	return func(app *App) (*executable.TypedExecutable[T, P], error) {
		reg, err := GetActionRegistration[T, P]()(app)
		if err != nil {
			return nil, err
		}

		executableAction := executable.NewExecutableAction(app.modelConfig, reg.GetTypeDefinition())
		act, ok := any(executableAction.Action).(*T)
		if !ok {
			return nil, fmt.Errorf("could nt ")
		}

		return &executable.TypedExecutable[T, P]{
			ExecutableAction: executableAction,
			Action:           act,
		}, nil
	}
}

func GetTypedAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*executable.Action[T, P], error) {
	return func(app *App) (*executable.Action[T, P], error) {
		reg, err := GetActionRegistration[T, P]()(app)
		if err != nil {
			return nil, err
		}

		return executable.NewAction(app.modelConfig, *reg, props), nil
	}
}
