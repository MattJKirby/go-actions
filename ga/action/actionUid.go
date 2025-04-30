package action

import (
	"fmt"
	"go-actions/ga/app/config"
)

type ActionUid struct {
	actionConfig *ActionConfig
	actionName string
	uid string
}

func NewActionUid(config *config.GlobalConfig, actionConfig *ActionConfig, actionName string) *ActionUid {
	return &ActionUid{
		actionConfig: actionConfig,
		actionName: actionName,
		uid: config.UidGenerator.GenerateUid(),
	}
}


func (au *ActionUid) GetUid() string {
	return fmt.Sprintf(au.actionConfig.UidFormat, au.actionName, au.uid)
}