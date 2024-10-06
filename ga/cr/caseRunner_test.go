package cr

import "testing"

func TestCaseDefaults( t *testing.T){
	defaultTestCase := TestCase[int, bool]{}
	exepcted := TestCase[int, bool]{
		Name: "",
		Input: 0,
		Expected: false,
		Description: "",
		Error: false,
	}

	t.Run("case defaults", func(t *testing.T) {
		if defaultTestCase != exepcted {
			t.Errorf("expected %v, got %v", exepcted, defaultTestCase)
		}
	})
}


func TestCaseRunner(t *testing.T) {

	isEven := func (n int) bool {
		return n%2 == 0
	}

	tests := []TestCase[int, bool]{
		{Name: "tc1", Input: 1, Expected: false},
		{Name: "tc2", Input: 2, Expected: true},
	}

	assert := func(test TestCase[int, bool]) {
		actual := isEven(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: exepcted %t, got %t", test.Name, actual, test.Expected)
		}
	}

	CaseRunner(t, tests, assert)
}
