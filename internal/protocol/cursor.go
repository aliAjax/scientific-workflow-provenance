package protocol

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Cursor struct {
	Created int64
	ID      string
}

func Encode(c Cursor) string { b, _ := json.Marshal(c); return base64.RawURLEncoding.EncodeToString(b) }
func Decode(v string) (Cursor, error) {
	b, e := base64.RawURLEncoding.DecodeString(v)
	if e != nil {
		return Cursor{}, e
	}
	var c Cursor
	if e = json.Unmarshal(b, &c); e != nil {
		return Cursor{}, e
	}
	if c.ID == "" {
		return Cursor{}, fmt.Errorf("cursor id missing")
	}
	return c, nil
}
func NextID(prefix string, n int) string { return prefix + "-" + strconv.Itoa(n) }
func Checksum(v string) string           { h := sha256.Sum256([]byte(v)); return fmt.Sprintf("%x", h[:]) }
func Normalize(v string) string          { return strings.TrimSpace(strings.ToLower(v)) }
