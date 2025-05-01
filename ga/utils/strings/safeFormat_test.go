package strings

import (
	"go-actions/ga/utils/testing/assert"
	"testing"
)

func TestValidatePlaceholders(t *testing.T) {
	tests := []struct {
		name             string
		format           string
		restrictToString bool
		expectedCount    int
		expectError      bool
	}{
		{
			name:             "valid %s only",
			format:           "action:%s:uid:%s",
			restrictToString: true,
			expectedCount:    2,
			expectError:      false,
		},
		{
			name:             "mixed %s and %d with restriction",
			format:           "action:%s:uid:%d",
			restrictToString: true,
			expectedCount:    0,
			expectError:      true,
		},
		{
			name:             "escaped percent signs",
			format:           "%%s %%d",
			restrictToString: true,
			expectedCount:    0,
			expectError:      false,
		},
		{
			name:             "non-restricted %d",
			format:           "id:%d",
			restrictToString: false,
			expectedCount:    1,
			expectError:      false,
		},
		{
			name:             "trailing percent sign",
			format:           "invalid %",
			restrictToString: true,
			expectedCount:    0,
			expectError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count, err := validatePlaceholders(tt.format, tt.restrictToString)
			assert.Equals(t, tt.expectedCount, count)
			assert.Equals(t, tt.expectError, err != nil)
		})
	}
}

func TestSafeFormat(t *testing.T) {
	tests := []struct {
		name        string
		format      string
		values      []string
		expected    string
		expectError bool
	}{
		{
			name:        "valid format with 2 placeholders",
			format:      "resource:%s:type:%s",
			values:      []string{"foo", "bar"},
			expected:    "resource:foo:type:bar",
			expectError: false,
		},
		{
			name:        "invalid format with %d",
			format:      "id:%d",
			values:      []string{"123"},
			expected:    "",
			expectError: true,
		},
		{
			name:        "too few values",
			format:      "uid:%s:name:%s",
			values:      []string{"only-one"},
			expected:    "",
			expectError: true,
		},
		{
			name:        "too many values",
			format:      "uid:%s",
			values:      []string{"val1", "val2"},
			expected:    "",
			expectError: true,
		},
		{
			name:        "escaped percent signs",
			format:      "progress:%% %s",
			values:      []string{"done"},
			expected:    "progress:% done",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeFormat(tt.format, tt.values...)
			assert.Equals(t, tt.expected, result)
			assert.Equals(t, tt.expectError, err != nil)
		})
	}
}
