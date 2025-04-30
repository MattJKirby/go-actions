package uid

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestWithNamespace(t *testing.T) {
	uid := NewResourceUid(mockGlobalConfig, WithNamespace("TestNameSpace"))
	assert.Equals(t, "TestNameSpace", uid.namespace)
}

func TestWithResource(t *testing.T) {
	uid := NewResourceUid(mockGlobalConfig, WithResource("Resource"))
	assert.Equals(t, "Resource", uid.resource)
}