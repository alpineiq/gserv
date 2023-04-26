package internal

import (
	"encoding/json"

	"github.com/alpineiq/otk"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalIndent(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "\t")
}

func Unmarshal(buf []byte, val any) error {
	return json.Unmarshal(buf, val)
}

func UnmarshalString(buf string, val any) error {
	return json.Unmarshal(otk.UnsafeBytes(buf), val)
}
