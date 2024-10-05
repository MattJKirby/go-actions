package cr

import (
	"fmt"
	"testing"
)

type TestCase[input any, expected any] struct {
	Name     string
	Input    input
	Expected expected
}

type AssertFn[input any, expected any] func(test TestCase[input, expected])

func CaseRunner[input any, expected any](t *testing.T, cases []TestCase[input, expected], assertFn AssertFn[input, expected]) {
	for caseNum, testCase := range cases {
		caseName := fmt.Sprintf("Case %d: '%s'", caseNum, testCase.Name)

		t.Run(caseName, func(t *testing.T) {
			assertFn(testCase)
		})
	}
}
