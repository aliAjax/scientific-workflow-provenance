package transport

import (
	"net/http"
	"scientific-workflow-provenance/pkg/pagination"
	"strconv"
)

func page(r *http.Request) pagination.Request {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	return pagination.Parse(limit, offset)
}
func pageHeaders(w http.ResponseWriter, total int, next bool) {
	w.Header().Set("X-Total-Count", strconv.Itoa(total))
	if next {
		w.Header().Set("X-Next-Page", "true")
	}
}
