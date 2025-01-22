package utils

import (
	"fmt"
	"strconv"
)

const (
	defaultPageSize = 25
)

type PaginationQuery struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginationParams struct {
	Limit  int
	Offset int
}

type Pagination struct {
	PaginationQuery
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_page"`
}

// Based on page and page size fetch desired limit and offset with this function.
// Applied validations, so if anything goes wrong then return error in that case.
func GetPaginationParameters(pageStr string, pageSizeStr string) (paginationParams PaginationParams, paginationQuery PaginationQuery, err error) {
	paginationQuery.Page, err = strconv.Atoi(pageStr)
	if pageStr != "" && err != nil {
		return paginationParams, paginationQuery, fmt.Errorf("can not parse page")
	}

	paginationQuery.PageSize, err = strconv.Atoi(pageSizeStr)
	if pageSizeStr != "" && err != nil {
		return paginationParams, paginationQuery, fmt.Errorf("can not parse pageSize")
	}

	if paginationQuery.Page < 0 || paginationQuery.PageSize < 0 {
		return paginationParams, paginationQuery, fmt.Errorf("invalid pagination parameters provided")
	}

	// use default values, if not provided by default
	if paginationQuery.PageSize == 0 {
		paginationQuery.PageSize = defaultPageSize
	}

	// calculate limit and offset based on page and page size
	paginationParams.Limit = paginationQuery.PageSize
	paginationParams.Offset = paginationQuery.Page * paginationQuery.PageSize

	return paginationParams, paginationQuery, nil
}

func GetPaginatedResponse(paginationQuery PaginationQuery, totalRecords int) Pagination {
	totalPages := (totalRecords + paginationQuery.PageSize - 1) / paginationQuery.PageSize
	return Pagination{
		PaginationQuery: paginationQuery,
		TotalRecords:    totalRecords,
		TotalPages:      totalPages,
	}
}
