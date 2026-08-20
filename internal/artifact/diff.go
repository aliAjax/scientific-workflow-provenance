package artifact

import (
	"fmt"
	"reflect"
)

type Change struct {
	Field         string
	Before, After any
}

func Diff(a, b Artifact) []Change {
	out := []Change{}
	if a.Name != b.Name {
		out = append(out, Change{"name", a.Name, b.Name})
	}
	if a.Schema != b.Schema {
		out = append(out, Change{"schema", a.Schema, b.Schema})
	}
	if a.Unit != b.Unit {
		out = append(out, Change{"unit", a.Unit, b.Unit})
	}
	if a.Value != b.Value {
		out = append(out, Change{"value", a.Value, b.Value})
	}
	if a.Status != b.Status {
		out = append(out, Change{"status", a.Status, b.Status})
	}
	return out
}
func Equal(a, b Artifact) bool { return reflect.DeepEqual(a, b) }
func RequireDigest(a Artifact) error {
	if len(a.Digest) < 16 {
		return fmt.Errorf("digest too short")
	}
	return nil
}
