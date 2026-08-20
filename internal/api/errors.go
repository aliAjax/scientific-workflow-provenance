package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func WrapCause(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("api error: %v", err)
}

func IsWrapped(err, target error) bool { return err != nil && target != nil && errors.Is(err, target) }

type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId,omitempty"`
}

func WriteError(w http.ResponseWriter, status int, code, message, id string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": Error{code, message, id}})
}
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
