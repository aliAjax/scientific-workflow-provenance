package expr

import (
	"fmt"
	"strconv"
	"strings"
)

type Op string

const (
	Eq  Op = "=="
	Ne  Op = "!="
	Gt  Op = ">"
	Gte Op = ">="
	Lt  Op = "<"
	Lte Op = "<="
)

type Predicate struct {
	Field string
	Op    Op
	Value float64
}

func Parse(input string) (Predicate, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return Predicate{}, fmt.Errorf("expression must be field op value")
	}
	v, e := strconv.ParseFloat(parts[2], 64)
	if e != nil {
		return Predicate{}, e
	}
	return Predicate{Field: parts[0], Op: Op(parts[1]), Value: v}, nil
}
func (p Predicate) Eval(fields map[string]float64) bool {
	v, ok := fields[p.Field]
	if !ok {
		return false
	}
	switch p.Op {
	case Eq:
		return v == p.Value
	case Ne:
		return v != p.Value
	case Gt:
		return v > p.Value
	case Gte:
		return v >= p.Value
	case Lt:
		return v < p.Value
	case Lte:
		return v <= p.Value
	}
	return false
}
