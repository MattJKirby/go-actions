package references

import (
	"fmt"
	"go-actions/ga/action/model"
)

type ActionReference struct {
	ReferenceUid    string `json:"ReferenceUid"`
	SourceActionUid string `json:"SourceActionUid"`
	TargetActionUid string `json:"TargetActionUid"`
}

func NewActionReference(config model.ActionConfig, sourceUid string, targetUid string) *ActionReference {
	return &ActionReference{
		ReferenceUid:    fmt.Sprintf("ref:%s", config.GenerateUid()),
		SourceActionUid: sourceUid,
		TargetActionUid: targetUid,
	}
}
