package provenance

import (
	"encoding/json"
	"fmt"
)

func (g *Graph) MarshalJSON() ([]byte, error) { return json.Marshal(g.Export()) }
func ValidateChain(data map[string]any) error {
	head, ok := data["head"].(string)
	if !ok || head == "" {
		return fmt.Errorf("provenance head missing")
	}
	return nil
}
