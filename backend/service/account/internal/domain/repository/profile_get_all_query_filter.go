package repository

import (
	"strings"
	"time"
)

type ProfileGetAllQueryFilter struct {
	Username   string `query:"user_name"`
	Date       string `query:"date"`
	IsArchived bool   `query:"is_archived"`
}

func DefaultAccountGetAllQueryFilter() *ProfileGetAllQueryFilter {
	return &ProfileGetAllQueryFilter{
		Username:   "",
		Date:       "",
		IsArchived: false,
	}
}

func NewAccountGetAllQueryFilter(m map[string]any) (*ProfileGetAllQueryFilter, error) {
	var filter ProfileGetAllQueryFilter = *DefaultAccountGetAllQueryFilter()
	if username := m["user_name"]; username != nil && strings.Trim(username.(string), " ") != "" {
		filter.Username = username.(string)
	}

	if dt, exists := m["date"]; exists && dt != nil && strings.Trim(dt.(string), " ") != "" {
		filter.Date = dt.(time.Time).String()
	}

	if isArchived, exists := m["is_archived"]; exists && isArchived != nil {
		filter.IsArchived = isArchived.(bool)
	}

	return &filter, nil
}
