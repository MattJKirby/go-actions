package action

import (
	"go-actions/ga/action/model"
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
)

type ActionInstance struct {
	Uid   uid.ResourceUid    `json:"uid"`
	Name  string             `json:"name"`
	Model *model.ActionModel `json:"model"`
}

func NewActionInstance(globalConfig *config.GlobalConfig, typedef *TypeDefinition) *ActionInstance {
	uid := uid.NewUidBuilder().WithResource(typedef.TypeName).WithUid(globalConfig.UidProvider.New()).Build()
	return &ActionInstance{
		Uid:   uid,
		Name:  typedef.TypeName,
		Model: model.NewActionModel(globalConfig, uid),
	}
}

func (ai ActionInstance) GetId() string {
	return ai.Uid.FullUid()
}
