package provenance

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrInvalidChain = errors.New("invalid provenance chain")

func (g *Graph) MarshalJSON() ([]byte, error) { return json.Marshal(g.Export()) }
func ValidateChain(data map[string]any) error {
	head, ok := data["head"].(string)
	if !ok || head == "" {
		return fmt.Errorf("provenance head missing: %v", ErrInvalidChain)
	}
	return nil
}
