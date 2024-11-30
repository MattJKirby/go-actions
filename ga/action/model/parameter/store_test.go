package parameter

import (
	"encoding/json"
	"go-actions/ga/cr"
	"go-actions/ga/cr/asserts"
	"testing"
)

func TestGetOrDefault(t *testing.T) {
	store := NewStore()
	expectedGetParam := NewActionParameter("getParam", 0)
	expectedDefaultParam := NewActionParameter("defaultParam", 0)

	t.Run("test Get path", func(t *testing.T) {
		param := GetOrDefault("getParam", 0)(store)
		asserts.Equals(t, expectedGetParam, param)
	})

	t.Run("test Default path", func(t *testing.T) {
		param := GetOrDefault("defaultParam", 0)(store)
		asserts.Equals(t, expectedDefaultParam, param)
		asserts.Equals(t, &expectedDefaultParam, &param)
	})
}

func TestCustomMarshal(t *testing.T) {
	store := NewStore()
	GetOrDefault("param", 0)(store)
	expectedJson := `{"param":{"name":"param","value":0}}`

	t.Run("test custom json marshal", func(t *testing.T) {
		marshalled, _ := json.Marshal(store)
		asserts.Equals(t, expectedJson, string(marshalled))
	})
}

func TestCustomUnmarshal(t *testing.T) {

	tests := []cr.TestCase[string, int]{
		{Name: "test valid input", Input: `{"param":{"name":"param","value":100}}`, Expected: 100, Error: false},
		{Name: "test bad input", Input: "0", Expected: 0, Error: true},
		{Name: "test bad parameter name", Input: `{"bad name":{"name":"param","value":100}}`, Expected: 0, Error: true},
		{Name: "test bad parameter value name", Input: `{"param":{"name":"paramx","value":100}}`, Expected: 0, Error: true},
	}

	cr.CaseRunner(t, tests, func(test cr.TestCase[string, int]) {
		expectedParam := &ActionParameter[int]{name: "param", defaultValue: 0, value: test.Expected}

		store := NewStore()
		GetOrDefault("param", 0)(store)

		err := json.Unmarshal([]byte(test.Input), store)
		param := GetOrDefault("param", 0)(store)
		asserts.Equals(t, expectedParam, param)

		hasErr := err != nil
		if test.Error != hasErr {
			t.Errorf("error unmarshalling store: got %v", err)
		}
	})
}
