package utils_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"github.com/viswals/core/pkg/utils"
)

func TestApplyFilters(t *testing.T) {
	tests := []struct {
		name         string
		filters      []utils.Filter
		isCountQuery bool
		expectedSQL  string
		expectedArgs []interface{}
	}{
		{
			name: "Equal filter",
			filters: []utils.Filter{
				{Field: "name", Operator: "eq", Value: "John"},
			},
			isCountQuery: false,
			expectedSQL:  "SELECT * FROM users WHERE name = ?",
			expectedArgs: []interface{}{"John"},
		},
		{
			name: "Like filter",
			filters: []utils.Filter{
				{Field: "name", Operator: "like", Value: "Doe"},
			},
			isCountQuery: false,
			expectedSQL:  "SELECT * FROM users WHERE name LIKE ?",
			expectedArgs: []interface{}{"%Doe%"},
		},
		{
			name: "Greater than or equal filter",
			filters: []utils.Filter{
				{Field: "age", Operator: "gte", Value: 30},
			},
			isCountQuery: false,
			expectedSQL:  "SELECT * FROM users WHERE age >= ?",
			expectedArgs: []interface{}{30},
		},
		{
			name: "Sorting",
			filters: []utils.Filter{
				{Field: "name", Operator: "eq", Value: "John"},
				{Field: "age", Sort: true, Order: "DESC"},
			},
			isCountQuery: false,
			expectedSQL:  "SELECT * FROM users WHERE name = ? ORDER BY age DESC",
			expectedArgs: []interface{}{"John"},
		},
		{
			name: "Count query ignores sorting",
			filters: []utils.Filter{
				{Field: "age", Operator: "gte", Value: 30},
				{Field: "age", Sort: true, Order: "ASC"},
			},
			isCountQuery: true,
			expectedSQL:  "SELECT * FROM users WHERE age >= ?",
			expectedArgs: []interface{}{30},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			baseQuery := sq.Select("*").From("users")
			query := utils.ApplyFilters(baseQuery, test.filters, test.isCountQuery)

			sql, args, err := query.ToSql()
			assert.NoError(t, err)
			assert.Equal(t, test.expectedSQL, sql)
			assert.Equal(t, test.expectedArgs, args)
		})
	}
}
