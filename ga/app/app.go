package app

import (
	"context"
	"go-actions/ga/action"
	"go-actions/ga/action/executable"
	"go-actions/ga/app/registration"
	"go-actions/ga/utils"
	"reflect"
)

type App struct {
	ctx            context.Context
	actionRegistry *registration.ActionRegistry
}

func NewApp() *App {
	return &App{
		ctx:            context.Background(),
		actionRegistry: registration.NewActionRegistry(),
	}
}

func RegisterAction[T action.GoAction, P action.GoActionProps](reg *action.GoActionRegistration[T, P]) func(*App) *registration.RegisteredAction[T, P] {
	return func(app *App) *registration.RegisteredAction[T, P] {
		action, _ := registration.NewRegisteredAction(reg)
		return registration.AcceptAction(action)(app.actionRegistry)
	}
}

func GetActionRegistration[T action.GoAction, P action.GoActionProps](action action.GoAction) func(*App) (*registration.RegisteredAction[T, P], error) {
	return func(app *App) (*registration.RegisteredAction[T, P], error) {
		actionType := utils.GetValueType(reflect.TypeOf(action))
		return registration.GetAction[T, P](actionType)(app.actionRegistry)
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](a action.GoAction) func(*App) (*executable.Action[T, P], error) {
	return func(app *App) (*executable.Action[T, P], error) {
		reg, err := GetActionRegistration[T, P](a)(app)
		if err != nil {
			return nil, err
		}
		return executable.NewAction[T, P](*reg), nil
	}
}
