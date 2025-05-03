package uid

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestWithNamespace(t *testing.T) {
	uid := NewResourceUid(WithNamespace("TestNameSpace"))
	assert.Equals(t, "TestNameSpace", uid.namespace)
}

func TestWithResource(t *testing.T) {
	uid := NewResourceUid(WithResource("Resource"))
	assert.Equals(t, "Resource", uid.resource)
}

func TestWithUid(t *testing.T) {
	uid := NewResourceUid(WithUid("Uid"))
	assert.Equals(t, "Uid", uid.uid)
}

func TestWithSubResource(t *testing.T) {
	uid := NewResourceUid(WithSubResource("subResource"))
	assert.Equals(t, "subResource", uid.subResourceType)
}

func TestWithSubResourceId(t *testing.T) {
	uid := NewResourceUid(WithSubResourceId("subResourceId"))
	assert.Equals(t, "subResourceId", uid.subResourceId)
}
