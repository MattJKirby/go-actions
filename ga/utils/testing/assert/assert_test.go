package assert

import "testing"

type testStruct struct {
	val  int
	flag bool
}

func TestAssertEquals(t *testing.T) {
	a := 1
	b := 1
	t.Run("test int equality", func(t *testing.T) {
		Equals(t, a, b)
	})

	x := testStruct{}
	y := testStruct{}

	t.Run("test struct equality", func(t *testing.T) {
		Equals(t, x, y)
	})
}
