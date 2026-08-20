package observability

import "context"

type key string

const requestKey key = "request-id"

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestKey, id)
}
func RequestID(ctx context.Context) string { v, _ := ctx.Value(requestKey).(string); return v }
func Fields(ctx context.Context) map[string]any {
	if id := RequestID(ctx); id != "" {
		return map[string]any{"request_id": id}
	}
	return map[string]any{}
}
