package sample

import (
	"fmt"
	"strings"
)

func Validate(v Sample) error {
	if strings.TrimSpace(v.ID) == "" {
		return fmt.Errorf("sample id required")
	}
	if strings.TrimSpace(v.Source) == "" {
		return fmt.Errorf("sample source required")
	}
	if v.State != Active && v.State != Frozen && v.State != Destroyed {
		return fmt.Errorf("invalid state %s", v.State)
	}
	return nil
}
func CanProcess(v Sample) bool { return v.State == Active && len(v.ParentIDs) < 100 }
