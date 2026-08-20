package quality

import (
	"math"
	"sort"
)

type Summary struct {
	Count                  int
	Mean, Min, Max, StdDev float64
	Missing                int
}

func Summarize(values []float64) Summary {
	if len(values) == 0 {
		return Summary{}
	}
	s := Summary{Count: len(values), Min: values[0], Max: values[0]}
	for _, v := range values {
		if math.IsNaN(v) {
			s.Missing++
			continue
		}
		s.Mean += v
		if v < s.Min {
			s.Min = v
		}
		if v > s.Max {
			s.Max = v
		}
	}
	valid := s.Count - s.Missing
	if valid > 0 {
		s.Mean /= float64(valid)
		for _, v := range values {
			if !math.IsNaN(v) {
				s.StdDev += (v - s.Mean) * (v - s.Mean)
			}
		}
		s.StdDev = math.Sqrt(s.StdDev / float64(valid))
	}
	return s
}
func Percentile(values []float64, p float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	v := append([]float64(nil), values...)
	sort.Float64s(v)
	if p <= 0 {
		return v[0]
	}
	if p >= 1 {
		return v[len(v)-1]
	}
	idx := p * float64(len(v)-1)
	lo := int(idx)
	hi := lo + 1
	if hi >= len(v) {
		return v[lo]
	}
	return v[lo] + (v[hi]-v[lo])*(idx-float64(lo))
}
