package model

import (
	"fmt"
	"go-actions/ga/action/model/references"
	"go-actions/ga/action/model/store"
)

type ActionModel struct {
	ActionName string                          `json:"name"`
	ActionUid  string                          `json:"uid"`
	Parameters *store.PropertyStore[any]       `json:"parameters"`
	Inputs     *store.ActionPropertyStore[references.ActionInput]  `json:"inputs"`
	Outputs    *store.ActionPropertyStore[references.ActionOutput] `json:"outputs"`
}

type ActionConfig interface {
	GenerateUid() string
}

func NewActionModel(typename string, config ActionConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, config.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: store.NewPropertyStore[any](),
		Inputs:     store.NewActionPropertyStore[references.ActionInput](false),
		Outputs:    store.NewActionPropertyStore[references.ActionOutput](false),
	}
}
