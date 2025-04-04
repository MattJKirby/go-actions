package marshalling

import (
	"go-actions/ga/utils/testing/assert"

	"testing"
)

func TestStrictDecode(t *testing.T) {
	type decodeType struct {
		A string
		B string
	}

	tests := []struct {
		name     string
		input    string
		expected *decodeType
		err      bool
	}{
		{name: "valid", input: `{"A":"atest","B":"btest"}`, expected: &decodeType{"atest", "btest"}, err: false},
		{name: "unexpected value", input: `{"A":"","B":"","C":"ctest"}`, expected: &decodeType{}, err: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			decoded := &decodeType{}
			_, err := StrictDecode([]byte(test.input), decoded)
			hasErr := err != nil

			assert.Equals(t, test.err, hasErr)
			assert.Equals(t, test.expected, decoded)
		})
	}
}
