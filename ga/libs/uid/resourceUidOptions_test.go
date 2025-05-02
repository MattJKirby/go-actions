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

func TestWithSubResource(t *testing.T) {
	uid := NewResourceUid(mockGlobalConfig, WithSubResource("subResource"))
	assert.Equals(t, "subResource", uid.subResourceType)
}

func TestWithSubResourceId(t *testing.T) {
	uid := NewResourceUid(mockGlobalConfig, WithSubResourceId("subResourceId"))
	assert.Equals(t, "subResourceId", uid.subResourceId)
}
