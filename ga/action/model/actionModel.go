package model

import (
	"fmt"
	"go-actions/ga/action/model/input"
	"go-actions/ga/action/model/output"
	"go-actions/ga/action/model/store"
	"go-actions/ga/app/config"
)

type ActionModel struct {
	ActionName string                                           `json:"name"`
	ActionUid  string                                           `json:"uid"`
	Parameters *store.PropertyStore[store.IdentifiableProperty] `json:"parameters"`
	Inputs     *store.PropertyStore[input.ActionInput]          `json:"inputs"`
	Outputs    *store.PropertyStore[output.ActionOutput]        `json:"outputs"`
}

func NewActionModel(typename string, globalConfig *config.GlobalConfig) *ActionModel {
	ActionUid := fmt.Sprintf("%s:%s", typename, globalConfig.UidGenerator.GenerateUid())
	return &ActionModel{
		ActionName: typename,
		ActionUid:  ActionUid,
		Parameters: store.NewPropertyStore[store.IdentifiableProperty](false),
		Inputs:     store.NewPropertyStore[input.ActionInput](false),
		Outputs:    store.NewPropertyStore[output.ActionOutput](false),
	}
}
