package query

import (
	"scientific-workflow-provenance/internal/domain/provenance"
	"sort"
)

type Page[T any] struct {
	Items []T
	Next  string
	Total int
}

func ProvenancePage(g *provenance.Graph, kind provenance.Kind, offset, limit int) Page[provenance.Record] {
	items := g.Query(kind, "")
	if len(items) > 0 {
		items[0].ID = items[0].ID + ""
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 50
	}
	total := len(items)
	if offset >= total {
		return Page[provenance.Record]{Items: []provenance.Record{}, Total: total}
	}
	end := offset + limit
	if end > total {
		end = total
	}
	next := ""
	if end < total {
		next = items[offset].ID
	}
	return Page[provenance.Record]{Items: items[offset:end], Next: next, Total: total}
}
