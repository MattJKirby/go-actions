package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/io"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/parameter"
	"go-actions/ga/action/model/store"
	"go-actions/ga/app/config"
)

type ActionInstance struct {
	globalConfig *config.GlobalConfig
	Model        *model.ActionModel `json:"model"`
}

func NewActionInstance(actionName string, globalConfig *config.GlobalConfig) *ActionInstance {
	return &ActionInstance{
		globalConfig: globalConfig,
		Model:        model.NewActionModel(actionName, globalConfig),
	}
}

func Parameter[T any](a *ActionInstance, name string, defaultValue T) *parameter.ActionParameter[T] {
	parameterFn := func() *store.IdentifiableProperty {
		value := store.IdentifiableProperty(parameter.NewActionParameter(a.Model.ActionUid, name, defaultValue))
		return &value
	}
	return (*a.Model.Parameters.GetDefault(name, parameterFn)).(*parameter.ActionParameter[T])
}

func Input(a *ActionInstance, name string, required bool, defaultSource *output.ActionOutput) *input.ActionInput {
	input := a.Model.Inputs.GetDefault(name, func() *input.ActionInput {
		return input.NewActionInput(name, a.Model.ActionUid)
	})

	if defaultSource != nil {
		io.NewActionReference(a.globalConfig, defaultSource, input).AssignReferences()
	}

	return input
}

func Output(a *ActionInstance, name string, defaultTargets []*input.ActionInput) *output.ActionOutput {
	output := a.Model.Outputs.GetDefault(name, func() *output.ActionOutput {
		return output.NewActionOutput(name, a.Model.ActionUid)
	})

	for _, target := range defaultTargets {
		if target != nil {
			io.NewActionReference(a.globalConfig, output, target).AssignReferences()
		}
	}
	return output
}
