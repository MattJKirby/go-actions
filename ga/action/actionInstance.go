package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/app/config"
)

type ActionInstance struct {
	Uid string `json:"uid"`
	Model *model.ActionModel `json:"model"`
}

func NewActionInstance(actionName string, globalConfig *config.GlobalConfig, actionConfig *ActionConfig) *ActionInstance {
	return &ActionInstance{
		Model: model.NewActionModel(actionName, globalConfig),
	}
}

func (ai ActionInstance) GetResourceId() string {
	return ai.Model.ActionUid
}