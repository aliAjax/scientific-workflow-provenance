package worker

import "strings"

type Capability struct {
	Name     string
	Version  string
	Features []string
}

func (c Capability) Supports(feature string) bool {
	for _, v := range c.Features {
		if v == feature {
			return true
		}
	}
	return false
}
func (c Capability) Compatible(required []string) bool {
	for _, v := range required {
		if !c.Supports(v) {
			return false
		}
	}
	return strings.TrimSpace(c.Name) != ""
}
