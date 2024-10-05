package testing

import "testing"

func isEven(n int) bool {
	return n%2 == 0
}

func TestCaseRunner(t *testing.T){

	tests := []TestCase[int, bool]{
		{"tc1", 1, false},
		{"tc2", 2, true},
	}

	CaseRunner(t, tests, func(test TestCase[int, bool]) {
		actual := isEven(test.input)
		if actual != test.expected {
			t.Errorf("test %s: exepcted %t, got %t", test.name, actual, test.expected)
		}
	})
}