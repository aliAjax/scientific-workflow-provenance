package quality

import (
	"fmt"
	"scientific-workflow-provenance/pkg/units"
)

type Measurement struct {
	Metric string
	Value  float64
	Unit   string
}

func Convert(m Measurement, to string) (Measurement, error) {
	v, e := units.Convert(m.Value, m.Unit, to)
	if e != nil {
		return Measurement{}, fmt.Errorf("%s: %w", m.Metric, e)
	}
	m.Value = v
	m.Unit = to
	return m, nil
}
func Compare(a, b Measurement, tolerance float64) (bool, error) {
	x, e := units.Convert(b.Value, b.Unit, a.Unit)
	if e != nil {
		return false, e
	}
	return units.Close(a.Value, x, tolerance), nil
}
