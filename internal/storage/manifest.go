package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Manifest struct {
	RunID     string    `json:"runId"`
	Files     []string  `json:"files"`
	CreatedAt time.Time `json:"createdAt"`
	Digest    string    `json:"digest"`
}

func NewManifest(run string, files []string) Manifest {
	m := Manifest{RunID: run, Files: append([]string(nil), files...), CreatedAt: time.Now().UTC()}
	b, _ := json.Marshal(m)
	h := sha256.Sum256(b)
	m.Digest = hex.EncodeToString(h[:])
	return m
}
func (m Manifest) Valid() bool {
	if m.RunID == "" || m.Digest == "" {
		return false
	}
	return true
}
