package utils

import (
	"fmt"
)

const defaultLimit = 10

type PageFilter struct {
	Page  int
	Limit int
}

func (p *PageFilter) WithQuery(query string) string {
	var pageQ string
	var limitQ string

	if p.Limit == 0 || p.Page > 0 {
		p.Limit = defaultLimit
	}
	if p.Page > 0 {
		pageQ = fmt.Sprintf(" OFFSET %d", (p.Page-1)*p.Limit)
	}
	if p.Limit > 0 {
		limitQ = fmt.Sprintf(" LIMIT %d", p.Limit)
	}

	return fmt.Sprintf("%s%s%s", query, pageQ, limitQ)
}
