package sample

import (
	"math"
	"sort"
)

type MetricSet struct{ Values map[string][]float64 }

func NewMetricSet() *MetricSet { return &MetricSet{Values: map[string][]float64{}} }
func (m *MetricSet) Add(name string, value float64) {
	if !math.IsNaN(value) {
		m.Values[name] = append(m.Values[name], value)
	}
}
func (m *MetricSet) Count(name string) int { return len(m.Values[name]) }
func (m *MetricSet) Mean(name string) float64 {
	v := m.Values[name]
	if len(v) == 0 {
		return 0
	}
	s := 0.0
	for _, x := range v {
		s += x
	}
	return s / float64(len(v))
}
func (m *MetricSet) Min(name string) float64 {
	v := m.Values[name]
	if len(v) == 0 {
		return 0
	}
	x := v[0]
	for _, n := range v[1:] {
		if n < x {
			x = n
		}
	}
	return x
}
func (m *MetricSet) Max(name string) float64 {
	v := m.Values[name]
	if len(v) == 0 {
		return 0
	}
	x := v[0]
	for _, n := range v[1:] {
		if n > x {
			x = n
		}
	}
	return x
}
func (m *MetricSet) ValuesSorted(name string) []float64 {
	o := append([]float64(nil), m.Values[name]...)
	sort.Float64s(o)
	return o
}

func (m *MetricSet) ValuesCopy(name string) []float64 {
	return append([]float64(nil), m.Values[name]...)
}
func (m *MetricSet) Names() []string {
	o := []string{}
	for k := range m.Values {
		o = append(o, k)
	}
	sort.Strings(o)
	return o
}
