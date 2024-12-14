package marshalling

import (
	"encoding/json"
	"strings"
)

func StrictDecode[T any](data []byte, v *T) (*json.Decoder, error) {
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(v)
	if err != nil {
		return nil, err
	}

	return decoder, nil
}
