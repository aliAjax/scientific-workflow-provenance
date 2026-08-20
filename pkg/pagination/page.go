package pagination

type Request struct{ Limit, Offset int }

func Parse(limit, offset int) Request {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	return Request{limit, offset}
}
func Slice[T any](items []T, r Request) ([]T, bool) {
	if r.Offset >= len(items) {
		return []T{}, false
	}
	end := r.Offset + r.Limit
	if end > len(items) {
		end = len(items)
	}
	return items[r.Offset:end], end < len(items)
}
