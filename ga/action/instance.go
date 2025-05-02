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

func NewActionInstance(globalConfig *config.GlobalConfig, actionConfig *ActionConfig, typedef *TypeDefinition) *ActionInstance {
	return &ActionInstance{
		Uid:   uid.NewResourceUid(globalConfig, uid.WithResource(typedef.TypeName)),
		Model: model.NewActionModel(typedef.TypeName, globalConfig),
	}
}

func (ai ActionInstance) GetResourceId() string {
	return ai.Uid.GetString()
}
