package uid

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestGetString(t *testing.T) {
	uid := &ResourceUid{prefix: "ga", namespace: "core", resource: "someaction", uid: "uid", subResource: "sub", subResourceId: "subId"}
	assert.Equals(t, "ga:core:someaction:uid:sub:subid", uid.FullUid())
	assert.Equals(t, "ga:core:someaction:uid::", uid.BaseUid())
}

func TestMarshal(t *testing.T) {
	uid := &ResourceUid{prefix: "ga", namespace: "core", resource: "someaction", uid: "uid", subResource: "sub", subResourceId: "subId"}
	marshalled, err := json.Marshal(uid)

	assert.Equals(t, err, nil)
	assert.Equals(t, fmt.Sprintf(`"%s"`, uid.FullUid()), string(marshalled))
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
			jsonInput:        `"ga:mynamespace:resource:somevaliduid::"`,
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
			name:             "other resource",
			jsonInput:        `"ga:mynamespace:otherresource:uid::"`,
			expectedResource: "otherresource",
			expectedUid:      "uid",
			expectErr:        false,
		},
		{
			name:      "wrong format",
			jsonInput: `"ga:mynamespace:onlythreeparts"`,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uid := &ResourceUid{prefix: "ga", namespace: "myNamespace"}
			err := json.Unmarshal([]byte(test.jsonInput), uid)

			fmt.Println(uid)

			assert.Equals(t, test.expectErr, err != nil)
			assert.Equals(t, "myNamespace", uid.namespace)
			assert.Equals(t, test.expectedResource, uid.resource)
			assert.Equals(t, test.expectedUid, uid.uid)
		})
	}
}
