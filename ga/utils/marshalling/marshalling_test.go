package marshalling

import (
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestStrictDecode(t *testing.T) {
	type decodeType struct {
		A string
		B string
	}

	tests := []cr.TestCase[string, *decodeType]{
		{Name: "valid", Input: `{"A":"atest","B":"btest"}`, Expected: &decodeType{"atest","btest"}, Error: false},
		{Name: "unexpected value", Input: `{"A":"","B":"","C":"ctest"}`, Expected: &decodeType{}, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, *decodeType]) {
		decoded := &decodeType{}
		_, err := StrictDecode([]byte(test.Input), decoded)
		hasErr := err != nil

		asserts.Equals(t, test.Error, hasErr)
		asserts.Equals(t, test.Expected, decoded)
	})
}