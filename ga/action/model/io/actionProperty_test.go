package io

import (
	"go-actions/ga/testing/assert"
	"testing"
)

func TestGetPropertyId(t *testing.T) {
	prop := NewActionProperty("uid", "type", "name")
	assert.Equals(t, "uid:type:name", prop.GetPropertyId())
}
