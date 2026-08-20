package units

import (
	"fmt"
	"math"
)

type Converter func(float64) float64

var table = map[string]map[string]Converter{"g": {"mg": func(v float64) float64 { return v * 1000 }, "kg": func(v float64) float64 { return v / 1000 }}, "mg": {"g": func(v float64) float64 { return v / 1000 }}, "C": {"K": func(v float64) float64 { return v + 273.15 }}, "K": {"C": func(v float64) float64 { return v - 273.15 }}}

func Convert(value float64, from, to string) (float64, error) {
	if from == to {
		return value, nil
	}
	if m, ok := table[from]; ok {
		if f, ok := m[to]; ok {
			return f(value), nil
		}
	}
	return 0, fmt.Errorf("unsupported conversion %s -> %s", from, to)
}
func Close(a, b, tolerance float64) bool { return math.Abs(a-b) <= tolerance }
