package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/app/config"
)

type ActionInstance struct {
	Model        *model.ActionModel `json:"model"`
}

func NewActionInstance(actionName string, globalConfig *config.GlobalConfig) *ActionInstance {
	return &ActionInstance{
		Model:        model.NewActionModel(actionName, globalConfig),
	}
}

