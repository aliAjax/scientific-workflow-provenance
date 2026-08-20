package sample

import (
	"encoding/json"
	"fmt"
)

func Encode(v Sample) ([]byte, error) { return json.Marshal(v) }
func Decode(b []byte) (Sample, error) {
	var v Sample
	if e := json.Unmarshal(b, &v); e != nil {
		return Sample{}, e
	}
	if v.ID == "" {
		return Sample{}, fmt.Errorf("sample id required")
	}
	return v, nil
}
func Clone(v Sample) Sample { b, _ := json.Marshal(v); out, _ := Decode(b); return out }
