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

type PartialActionReference struct {
	ReferenceUid string `json:"ReferenceUid"`
	ActionUid    string `json:"ActionUid"`
}

func NewActionReference(config model.ActionConfig, sourceUid string, targetUid string) *ActionReference {
	return &ActionReference{
		ReferenceUid:    fmt.Sprintf("ref:%s", config.GenerateUid()),
		SourceActionUid: sourceUid,
		TargetActionUid: targetUid,
	}
}

func (ar ActionReference) GetId() string {
	return ar.ReferenceUid
}

func (ar *ActionReference) GetSourceReference() *PartialActionReference {
	return &PartialActionReference{
		ReferenceUid: ar.ReferenceUid,
		ActionUid:    ar.SourceActionUid,
	}
}

func (ar *ActionReference) GetTargetReference() *PartialActionReference {
	return &PartialActionReference{
		ReferenceUid: ar.ReferenceUid,
		ActionUid:    ar.TargetActionUid,
	}
}
