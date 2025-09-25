package repository

import (
	"strings"
)

type ProfileGetAllQueryFilter struct {
	Username string `query:"user_name"`
	// CreatedAt  []time.Time `query:"created_at"`
	IsArchived bool `query:"is_archived"`
}

func DefaultAccountGetAllQueryFilter() *ProfileGetAllQueryFilter {
	return &ProfileGetAllQueryFilter{
		Username: "",
		// Date:       "",
		IsArchived: false,
	}
}

func NewAccountGetAllQueryFilter(m map[string]any) (*ProfileGetAllQueryFilter, error) {
	var filter ProfileGetAllQueryFilter = *DefaultAccountGetAllQueryFilter()
	if username := m["user_name"]; username != nil && strings.Trim(username.(string), " ") != "" {
		filter.Username = username.(string)
	}

	// if dt, exists := m["date"]; exists && dt != nil && strings.Trim(dt.(string), " ") != "" {
	// 	filter.CreatedAt = time.P
	// }

	if isArchived, exists := m["is_archived"]; exists && isArchived != nil {
		filter.IsArchived = isArchived.(bool)
	}

	return &filter, nil
}
