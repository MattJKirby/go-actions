package io

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
)

type ActionReference struct {
	*uid.Reference
}


func NewActionReference(globalConfig *config.GlobalConfig, source *uid.ResourceUid, target *uid.ResourceUid) *ActionReference {
	return &ActionReference{
		Reference: uid.NewReference(source, target, uid.WithUid(globalConfig.UidGenerator.GenerateUid())),
	}
}

func (ar ActionReference) GetResourceId() string {
	return ar.Uid.GetUid()
}
