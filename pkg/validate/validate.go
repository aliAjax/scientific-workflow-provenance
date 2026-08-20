package validate

import (
	"fmt"
	"regexp"
	"strings"
)

var nameRx = regexp.MustCompile(`^[a-z][a-z0-9_.-]{1,62}$`)

func Name(v string) error {
	if !nameRx.MatchString(v) {
		return fmt.Errorf("invalid name %q", v)
	}
	return nil
}
func Required(fields ...string) error {
	for i, v := range fields {
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("field %d is required", i+1)
		}
	}
	return nil
}
func OneOf(value string, choices ...string) error {
	for _, c := range choices {
		if value == c {
			return nil
		}
	}
	return fmt.Errorf("%q is not an allowed value", value)
}
