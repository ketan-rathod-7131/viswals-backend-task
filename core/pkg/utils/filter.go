package utils

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type FilterOperator string

const (
	FilterOperatorEq     FilterOperator = "eq"
	FilterOperatorLike   FilterOperator = "like"
	FilterOperatorNeq    FilterOperator = "neq"
	FilterOperatorGte    FilterOperator = "gte"
	FilterOperatorLte    FilterOperator = "lte"
	FilterOperatorGt     FilterOperator = "gt"
	FilterOperatorLt     FilterOperator = "lt"
	FilterOperatorLimit  FilterOperator = "limit"
	FilterOperatorOffset FilterOperator = "offset"
)

// Filter is a helper struct that allows you to filter a database field.
type Filter struct {
	Field    string
	Operator FilterOperator
	Value    interface{}
	Sort     bool
	Order    string
}

// ApplyFilters applies filters to the base squirrel query.
func ApplyFilters(baseQuery sq.SelectBuilder, filters []Filter, isCountQuery bool) sq.SelectBuilder {

	// iterate all the filters and extend base query
	for _, filter := range filters {
		switch filter.Operator {
		case FilterOperatorEq:
			baseQuery = baseQuery.Where(sq.Eq{filter.Field: filter.Value})
		case FilterOperatorLike:
			baseQuery = baseQuery.Where(sq.Like{filter.Field: fmt.Sprintf("%%%s%%", filter.Value)})
		case FilterOperatorGte:
			baseQuery = baseQuery.Where(sq.GtOrEq{filter.Field: filter.Value})
		case FilterOperatorLte:
			baseQuery = baseQuery.Where(sq.LtOrEq{filter.Field: filter.Value})
		case FilterOperatorNeq:
			baseQuery = baseQuery.Where(sq.NotEq{filter.Field: filter.Value})
		case FilterOperatorLimit:
			baseQuery = baseQuery.Limit(uint64(filter.Value.(int)))
		case FilterOperatorOffset:
			baseQuery = baseQuery.Offset(uint64(filter.Value.(int)))
		}

		// for COUNT(*) queries, ignore sort field as aggregation is required for such queries ( COUNT, SUM etc. ).
		if !isCountQuery && filter.Sort {
			order := "ASC" // by default ascending order
			if filter.Order == "DESC" {
				order = "DESC"
			}

			baseQuery = baseQuery.OrderBy(fmt.Sprintf("%s %s", filter.Field, order))
		}
	}

	return baseQuery
}
