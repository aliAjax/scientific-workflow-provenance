package workflow

import (
	"encoding/json"
	"fmt"
)

func Encode(d Definition) ([]byte, error) { return json.MarshalIndent(d, "", "  ") }
func Decode(data []byte) (Definition, error) {
	var d Definition
	if err := json.Unmarshal(data, &d); err != nil {
		return Definition{}, err
	}
	if err := d.Validate(); err != nil {
		return Definition{}, fmt.Errorf("decode workflow: %w", err)
	}
	return d, nil
}
func Clone(d Definition) Definition                { b, _ := json.Marshal(d); v, _ := DecodeUnchecked(b); return v }
func DecodeUnchecked(b []byte) (Definition, error) { var d Definition; return d, json.Unmarshal(b, &d) }
