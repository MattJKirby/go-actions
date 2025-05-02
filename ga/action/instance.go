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
	uid := uid.NewResourceUid(globalConfig, uid.WithResource(typedef.TypeName))
	return &ActionInstance{
		Uid:   uid,
		Model: model.NewActionModel(uid, globalConfig),
	}
}

func (ai ActionInstance) GetResourceId() string {
	return ai.Uid.GetString()
}
