package io

import (
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestGetPropertyId(t *testing.T) {
	prop := NewActionProperty("uid", "type", "name")
	asserts.Equals(t, "uid:type:name", prop.GetPropertyId())
}
