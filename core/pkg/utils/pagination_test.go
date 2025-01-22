package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaginationParameters(t *testing.T) {
	tests := []struct {
		name           string
		pageStr        string
		pageSizeStr    string
		expectedLimit  int
		expectedOffset int
		expectedError  bool
	}{
		{
			name:           "valid input",
			pageStr:        "1",
			pageSizeStr:    "10",
			expectedLimit:  10,
			expectedOffset: 10,
			expectedError:  false,
		},
		{
			name:           "default page size",
			pageStr:        "2",
			pageSizeStr:    "",
			expectedLimit:  defaultPageSize,
			expectedOffset: 50,
			expectedError:  false,
		},
		{
			name:           "invalid page",
			pageStr:        "-1",
			pageSizeStr:    "10",
			expectedLimit:  0,
			expectedOffset: 0,
			expectedError:  true,
		},
		{
			name:           "invalid page size",
			pageStr:        "1",
			pageSizeStr:    "-5",
			expectedLimit:  0,
			expectedOffset: 0,
			expectedError:  true,
		},
		{
			name:           "non-numeric page",
			pageStr:        "abc",
			pageSizeStr:    "10",
			expectedLimit:  0,
			expectedOffset: 0,
			expectedError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			paginationParams, _, err := GetPaginationParameters(test.pageStr, test.pageSizeStr)

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedLimit, paginationParams.Limit)
				assert.Equal(t, test.expectedOffset, paginationParams.Offset)
			}
		})
	}
}

func TestGetPaginatedResponse(t *testing.T) {
	tests := []struct {
		name               string
		paginationQuery    PaginationQuery
		totalRecords       int
		expectedTotalPages int
	}{
		{
			name: "valid Pagination",
			paginationQuery: PaginationQuery{
				Page:     1,
				PageSize: 10,
			},
			totalRecords:       25,
			expectedTotalPages: 3,
		},
		{
			name: "exact Records",
			paginationQuery: PaginationQuery{
				Page:     0,
				PageSize: 10,
			},
			totalRecords:       20,
			expectedTotalPages: 2,
		},
		{
			name: "no Records",
			paginationQuery: PaginationQuery{
				Page:     0,
				PageSize: 10,
			},
			totalRecords:       0,
			expectedTotalPages: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pagination := GetPaginatedResponse(test.paginationQuery, test.totalRecords)
			assert.Equal(t, test.expectedTotalPages, pagination.TotalPages)
			assert.Equal(t, test.totalRecords, pagination.TotalRecords)
		})
	}
}
