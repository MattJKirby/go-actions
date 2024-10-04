package testing

import (
	"fmt"
	"testing"
)

type TestCase[input any, expected any] struct {
	name string
	input input
	expected expected
}

func CaseRunner[input any, expected any](t *testing.T, cases []TestCase[input,expected], assertion func (testcase TestCase[input, expected])){
	for caseNum, testCase := range cases {
		caseName := fmt.Sprintf("Case %d: '%s'", caseNum, testCase.name)
		
		t.Run(caseName, func(t *testing.T) {
			assertion(testCase)
		})
	}
}