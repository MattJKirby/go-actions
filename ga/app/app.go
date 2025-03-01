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

type InstantiatedTypedAction[T action.GoAction] struct {
	Instance *action.ActionInstance
	Action   *T
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

func GetActionByName(actionName string) func(*App) (*InitialisedAction, error) {
	return func(app *App) (*InitialisedAction, error) {
		typeDef, err := getRegisteredTypeDefinitionByName(actionName)(app.actionRegistry)
		if err != nil {
			return nil, err
		}
		return NewInitialisedAction(app, typeDef)
	}
}

func GetAction[T action.GoAction, P action.GoActionProps](props *P) func(*App) (*InstantiatedTypedAction[T], error) {
	return func(app *App) (*InstantiatedTypedAction[T], error) {
		reg, err := GetActionRegistration[T, P]()(app)
		if err != nil {
			return nil, err
		}

		// executableAction := executable.NewExecutableAction(app.modelConfig, reg.GetTypeDefinition())
		// act, ok := any(executableAction.Action).(*T)
		// if !ok {
		// 	return nil, fmt.Errorf("could nt ")
		// }

		instance := action.NewActionInstance(reg.TypeName, app.modelConfig)
		action := reg.Constructor(instance, *reg.DefaultProps)

		return &InstantiatedTypedAction[T]{
			Instance: instance,
			Action:   action,
		}, nil
	}
}
