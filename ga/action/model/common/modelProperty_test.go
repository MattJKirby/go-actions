package common

import (
	"go-actions/ga/libs/uid"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestGetPropertyId(t *testing.T) {
	modelUid := uid.NewResourceUid()
	prop := NewModelProperty(modelUid, "type", "name")
	assert.Equals(t, "ga:core:::type:name", prop.GetResourceId())
	assert.Equals(t, "ga:core::::", prop.GetActionUid())
}
