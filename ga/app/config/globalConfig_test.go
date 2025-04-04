package config

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

type mockUidGenerator struct {
	uid string
}
func (m *mockUidGenerator) GenerateUid() string {
	return m.uid
}

func TestWithCustomUidGenerator(t *testing.T) {
	config := NewGlobalConfig()
	mockGen := &mockUidGenerator{uid: "uid"}
	WithCustomUidGenerator(mockGen)(config)

	assert.Equals(t, "uid", config.UidGenerator.GenerateUid())
}