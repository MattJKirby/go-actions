package uid

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

// var mockUidGenerator = &testHelpers.MockUidGenerator{MockUid: "abc"}
// var mockGlobalConfig = &config.GlobalConfig{UidGenerator: mockUidGenerator}

func TestGetString(t *testing.T) {
	uid := NewResourceUid(WithResource("someAction"), WithUid("uid"))
	assert.Equals(t, "ga:core:someaction:uid::", uid.GetString())
}

func TestChildUid(t *testing.T) {
	parent := NewResourceUid(WithResource("parent"), WithUid("uid"))
	child := parent.FromParent(WithSubResource("subResource"))

	assert.Equals(t, "ga:core:parent:uid:subresource:", child.GetString())
}

func TestMarshal(t *testing.T) {
	uid := NewResourceUid(WithNamespace("testns"), WithResource("someAction"))

	marshalled, err := json.Marshal(uid)

	assert.Equals(t, err, nil)
	assert.Equals(t, fmt.Sprintf(`"%s"`, uid.GetString()), string(marshalled))
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name             string
		jsonInput        string
		expectedResource string
		expectedUid      string
		expectErr        bool
	}{
		{
			name:             "valid UID",
			jsonInput:        `"ga:mynamespace:myresource:somevaliduid::"`,
			expectedResource: "resource",
			expectedUid:      "somevaliduid",
			expectErr:        false,
		},
		{
			name:      "invalid prefix",
			jsonInput: `"wrongprefix:mynamespace:resource:uid::"`,
			expectErr: true,
		},
		{
			name:      "invalid namespace",
			jsonInput: `"ga:wrongnamespace:resource:uid::"`,
			expectErr: true,
		},
		{
			name:      "invalid resource",
			jsonInput: `"ga:mynamespace:wrongresource:uid::"`,
			expectErr: true,
		},
		{
			name:      "wrong format",
			jsonInput: `"ga:mynamespace:onlythreeparts"`,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uid := NewResourceUid(WithNamespace("myNamespace"), WithResource("myResource"), WithUid(""))
			marshalled := []byte(test.jsonInput)
			err := json.Unmarshal(marshalled, uid)

			assert.Equals(t, test.expectErr, err != nil)
			assert.Equals(t, "myNamespace", uid.namespace)
			assert.Equals(t, "myResource", uid.resource)
			assert.Equals(t, test.expectedUid, uid.uid)
		})
	}
}
