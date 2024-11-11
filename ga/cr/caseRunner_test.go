package cr

import (
	"fmt"
	"testing"
)

func TestCaseDefaults(t *testing.T) {
	defaultTestCase := TestCase[int, bool]{}
	exepcted := TestCase[int, bool]{
		Name:        "",
		Input:       0,
		Expected:    false,
		Description: "",
		Error:       false,
	}

	t.Run("case defaults", func(t *testing.T) {
		if defaultTestCase != exepcted {
			t.Errorf("expected %v, got %v", exepcted, defaultTestCase)
		}
	})
}

func exampleIsEven(n int) bool {
	return n%2 == 0
}

func TestCaseRunner(t *testing.T) {

	tests := []TestCase[int, bool]{
		{Name: "tc1", Input: 1, Expected: false},
		{Name: "tc2", Input: 2, Expected: true},
	}

	assert := func(test TestCase[int, bool]) {
		actual := exampleIsEven(test.Input)
		if actual != test.Expected {
			t.Errorf("test %s: exepcted %t, got %t", test.Name, actual, test.Expected)
		}
	}

	CaseRunner(t, tests, assert)
}

func exampleCanError(n int) (bool, error) {
	if n%3 == 0 {
		return false, fmt.Errorf("error")
	}
	return n%2 == 0, nil
}

func TestEqualityCaseRunner(t *testing.T) {

	tests := []TestCase[int, bool]{
		{Name: "tc1", Input: 1, Expected: false, Error: false},
		{Name: "tc2", Input: 2, Expected: true, Error: false},
		{Name: "tc3", Input: 2, Expected: true, Error: false},
		{Name: "tc4", Input: 3, Error: true},
	}

	EasyCaseRunner(t, tests, func(test TestCase[int, bool]) (bool, error) {
		return exampleCanError(test.Input)
	})
}
