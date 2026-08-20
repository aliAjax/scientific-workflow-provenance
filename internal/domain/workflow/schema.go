package workflow

import (
	"fmt"
	"regexp"
	"strings"
)

type Schema struct {
	Required []string          `json:"required"`
	Types    map[string]string `json:"types"`
}

func (s Schema) Validate(input map[string]any) error {
	for _, f := range s.Required {
		if _, ok := input[f]; !ok {
			return fmt.Errorf("required parameter missing: %s", f)
		}
	}
	for k, t := range s.Types {
		v, ok := input[k]
		if !ok {
			continue
		}
		if t == "string" {
			if _, ok := v.(string); !ok {
				return fmt.Errorf("%s must be string", k)
			}
		}
		if t == "number" {
			switch v.(type) {
			case float64, int, int64:
			default:
				return fmt.Errorf("%s must be number", k)
			}
		}
	}
	return nil
}
func NormalizeExpression(e string) string {
	return strings.Join(strings.Fields(regexp.MustCompile(`\s+`).ReplaceAllString(e, " ")), " ")
}
