package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/viswals/core/pkg/utils"
	"go.uber.org/zap"
)

func (c *Controller) GetAllUsers(g *gin.Context) {

	// pagination details
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")

	c.logger.Info("get all users", zap.String("page", pageStr), zap.String("page_size", pageSizeStr))

	// validate and parse pagination parameters
	paginationParams, paginationQuery, paginationError := utils.GetPaginationParameters(pageStr, pageSizeStr)
	if paginationError != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination parameters provided"})
		return
	}

	// create filters
	queryParams := g.Request.URL.Query()
	filters := make([]utils.Filter, 0)

	c.logger.Info("query parameters", zap.Any("query", queryParams))

	// sort filter
	sort := g.Query("sort")
	if arr := strings.Split(sort, ":"); len(arr) == 2 {
		order := "ASC"
		var field string = arr[0]

		if arr[1] == "DESC" {
			order = "DESC"
		}

		// allow sorting only for particular fields
		if field == "id" {
			filters = append(filters, utils.Filter{
				Field: field,
				Sort:  true,
				Order: order,
			})
		}
	}

	// filter by id min and max values
	idMin := queryParams.Get("id:min")
	idMax := queryParams.Get("id:max")

	if idMin != "" && idMax != "" {
		filters = append(filters, utils.Filter{
			Field:    "id",
			Operator: utils.FilterOperatorGte,
			Value:    idMin,
		})
		filters = append(filters, utils.Filter{
			Field:    "id",
			Operator: utils.FilterOperatorLte,
			Value:    idMax,
		})
	}

	c.logger.Debug("filters applied", zap.Any("filters", filters))

	// TODO: add other filters

	users, totalUsers, err := c.usecase.GetAllUsers(g, paginationParams, filters)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching users data"})
		return
	}

	// generate pagination response
	pagination := utils.GetPaginatedResponse(paginationQuery, totalUsers)

	g.JSON(http.StatusOK, gin.H{"data": users, "pagination": pagination})
}

func (c *Controller) GetUserById(g *gin.Context) {
	c.logger.Info("get user by id")

	// validations
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := c.usecase.GetUserById(g, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching user by id"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"data": user})
}
