package assert

import (
	"reflect"
	"testing"
)

func Equals[k any](t *testing.T, expected k, actual k) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("equals assertion failed: expected %v but got %v", expected, actual)
	}
}
