package flow

import (
	"go-actions/ga/action"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestNewFlowDefinition(t *testing.T){
	flowDef := NewFlowDefinition()
	expectedInstances := make(map[string]*action.ActionInstance)
	asserts.Equals(t, expectedInstances, flowDef.actions)
}