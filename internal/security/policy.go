package security

import (
	"errors"
	"net"
)

type Policy struct {
	AllowedNetworks  []*net.IPNet
	MaxArtifactBytes int64
	RequireReview    bool
}

func (p Policy) CheckNetwork(ip net.IP) error {
	for _, n := range p.AllowedNetworks {
		if n.Contains(ip) {
			return nil
		}
	}
	return errors.New("network not allowed")
}
func (p Policy) CheckSize(size int64) error {
	if p.MaxArtifactBytes > 0 && size > p.MaxArtifactBytes {
		return errors.New("artifact exceeds policy size")
	}
	return nil
}
func (p Policy) NeedsReview(risk string) bool {
	return p.RequireReview || risk == "high" || risk == "critical"
}
