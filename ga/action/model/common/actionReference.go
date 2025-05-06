package common

import (
	"go-actions/ga/app/config"
	"go-actions/ga/libs/uid"
)

type ActionReference struct {
	Uid uid.ResourceUid `json:"uid"`
	Source uid.ResourceUid `json:"source"`
	Target uid.ResourceUid `json:"target"`
}

func NewActionReference(globalConfig *config.GlobalConfig, source uid.ResourceUid, target uid.ResourceUid) *ActionReference {
	return &ActionReference{
		Uid: uid.NewUidBuilder().WithResource("Res").WithUid(globalConfig.UidGenerator.GenerateUid()).Build(),
		Source: source,
		Target: target,
	}
}

func (ar ActionReference) GetResourceId() string {
	return ar.Uid.FullUid()
}
