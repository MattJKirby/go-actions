package cr

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase[input any, expected any] struct {
	Name        string
	Description string
	Input       input
	Expected    expected
	Error       bool
}

type AssertFn[input any, expected any] func(test TestCase[input, expected])

func CaseRunner[input any, expected any](t *testing.T, cases []TestCase[input, expected], assertFn AssertFn[input, expected]) {
	t.Helper()
	for caseNum, testCase := range cases {
		caseName := fmt.Sprintf("Case %d: '%s'", caseNum, testCase.Name)

		t.Run(caseName, func(t *testing.T) {
			assertFn(testCase)
		})
	}
}

// Can't test error paths without mocking testing.T which is a bitch without another lib.
func EasyCaseRunner[in any, expect any](t *testing.T, cases []TestCase[in, expect], fn func(test TestCase[in, expect]) (expect, error)) {
	CaseRunner(t, cases, func(test TestCase[in, expect]) {
		actual, err := fn(test)

		if !test.Error && err != nil {
			t.Errorf("case '%s': unexpected error: %v", test.Name, err)
		}

		if test.Error && err == nil {
			t.Errorf("case '%s': expected error but got %v", test.Name, err)
		}

		if !reflect.DeepEqual(actual, test.Expected) {
			t.Errorf("case '%s': expected %v but got %v", test.Name, test.Expected, actual)
		}
	})
}
