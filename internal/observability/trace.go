package observability

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type traceKey struct{}

func StartSpan(ctx context.Context, name string) (context.Context, string) {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	id := hex.EncodeToString(b)
	return context.WithValue(ctx, traceKey{}, id), id
}
func TraceID(ctx context.Context) string { v, _ := ctx.Value(traceKey{}).(string); return v }
func SpanName(id, name string) string    { return fmt.Sprintf("%s:%s", id, name) }
