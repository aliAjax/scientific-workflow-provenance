package provenance

import (
	"fmt"
	"strings"
)

func ValidateRecord(r Record) error {
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("record id required")
	}
	if r.Kind != Entity && r.Kind != Activity && r.Kind != Agent {
		return fmt.Errorf("unknown record kind")
	}
	if r.Type == "" {
		return fmt.Errorf("record type required")
	}
	return nil
}
func ValidateRelation(r Relation) error {
	if r.From == "" || r.To == "" {
		return fmt.Errorf("relation endpoints required")
	}
	if r.From == r.To {
		return fmt.Errorf("self relation")
	}
	if r.Type == "" {
		return fmt.Errorf("relation type required")
	}
	return nil
}
