package uid

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	uid := NewUidBuilder().
		WithNamespace("Ns").
		WithResource("Res").
		WithUid("Uid").
		WithSubResource("Sub").
		WithSubResourceId("Id").
		Build()

	assert.Equals(t, "ga:ns:res:uid:sub:id", uid.FullUid())
	assert.Equals(t, "ga:ns:res:uid::", uid.BaseUid())
}
