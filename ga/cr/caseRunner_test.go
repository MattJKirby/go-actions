package cr

import "testing"

func isEven(n int) bool {
	return n%2 == 0
}

func TestCaseRunner(t *testing.T){

	tests := []TestCase[int, bool]{
		{"tc1", 1, false},
		{"tc2", 2, true},
	}

	assert := func(test TestCase[int, bool]){
		actual := isEven(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: exepcted %t, got %t", test.Name, actual, test.Expected)
		}
	}

	CaseRunner(t, tests, assert)
}