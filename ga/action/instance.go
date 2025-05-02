package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
)

type ActionInstance struct {
	Uid   *uid.ResourceUid   `json:"uid"`
	Model *model.ActionModel `json:"model"`
}

func NewActionInstance(actionName string, globalConfig *config.GlobalConfig, actionConfig *ActionConfig) *ActionInstance {
	return &ActionInstance{
		Uid:   uid.NewResourceUid(globalConfig, uid.WithResource(actionName)),
		Model: model.NewActionModel(actionName, globalConfig),
	}
}

func (ai ActionInstance) GetResourceId() string {
	return ai.Uid.GetString()
}
