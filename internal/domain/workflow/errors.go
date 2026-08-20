package workflow

import "fmt"

type ValidationError struct {
	Field  string
	Reason string
}

func (e ValidationError) Error() string { return fmt.Sprintf("%s: %s", e.Field, e.Reason) }

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}
	return e[0].Error()
}
func (e ValidationErrors) Add(field, reason string) ValidationErrors {
	return append(e, ValidationError{field, reason})
}
func (e ValidationErrors) HasErrors() bool { return len(e) > 0 }
