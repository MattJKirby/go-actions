package app

import (
	"context"
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
}

func NewApp(name string) *App {
	return &App{
		Name:           name,
		ctx:            context.Background(),
		actionRegistry: newActionRegistry(),
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

func GetTypedAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*executable.Action[T, P], error) {
	return func(app *App) (*executable.Action[T, P], error) {
		reg, err := GetActionRegistration[T, P]()(app)
		if err != nil {
			return nil, err
		}

		modelConfig := model.NewModelConfig()

		return executable.NewAction(modelConfig, *reg, props), nil
	}
}
