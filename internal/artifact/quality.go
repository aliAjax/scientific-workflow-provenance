package artifact

import (
	"fmt"
	"math"
	"scientific-workflow-provenance/pkg/units"
)

type Observation struct {
	Value  float64
	Unit   string
	Metric string
}

func Normalize(o Observation, target string) (Observation, error) {
	v, e := units.Convert(o.Value, o.Unit, target)
	if e != nil {
		return Observation{}, e
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return Observation{}, fmt.Errorf("non-finite observation")
	}
	o.Unit = target
	return o, nil
}
func Range(min, max float64) QualityGate { return QualityGate{Min: &min, Max: &max} }
