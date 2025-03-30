package model

import (
	"fmt"
	"go-actions/ga/action/model/references"
	"go-actions/ga/action/model/store"
)

type ActionModel struct {
	ActionName string                                           `json:"name"`
	ActionUid  string                                           `json:"uid"`
	Parameters *store.PropertyStore[store.IdentifiableProperty] `json:"parameters"`
	Inputs     *store.PropertyStore[references.ActionInput]     `json:"inputs"`
	Outputs    *store.PropertyStore[references.ActionOutput]    `json:"outputs"`
}

type ActionConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: store.NewPropertyStore[store.IdentifiableProperty](false),
		Inputs:     store.NewPropertyStore[references.ActionInput](false),
		Outputs:    store.NewPropertyStore[references.ActionOutput](false),
	}
}
