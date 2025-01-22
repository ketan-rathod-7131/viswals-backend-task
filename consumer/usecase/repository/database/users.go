package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/viswals/core/models"
	"github.com/viswals/core/pkg/utils"
	"go.uber.org/zap"
)

func (a *ConsumerDB) CreateUser(ctx context.Context, user models.User) (id string, err error) {

	query := `INSERT INTO users (email, firstname, lastname, parent_user_id, created_at, deleted_at, merged_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = a.DB.QueryRowContext(ctx, query, user.Email, user.FirstName, user.LastName, user.ParentUserId, user.CreatedAt, user.DeletedAt, user.MergedAt).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (g *ConsumerDB) GetUserByEmail(ctx context.Context, email string) (user models.User, err error) {

	query := "SELECT id, email, firstname, lastname, parent_user_id FROM users WHERE email = $1"
	row := g.DB.QueryRowContext(ctx, query, email)
	err = row.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.ParentUserId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (g *ConsumerDB) GetUserById(ctx context.Context, id int64) (user models.User, err error) {

	query := "SELECT id, email, firstname, lastname, parent_user_id, created_at, deleted_at, merged_at FROM users WHERE id = $1"
	row := g.DB.QueryRowContext(ctx, query, id)
	err = row.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.ParentUserId, &user.CreatedAt, &user.DeletedAt, &user.MergedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (g *ConsumerDB) GetAllUsers(ctx context.Context, pagination utils.PaginationParams, filters []utils.Filter) ([]models.User, int, error) {
	var users []models.User

	// build sql query
	baseQuery := sq.Select("id, email, firstname, lastname, parent_user_id, created_at, deleted_at, merged_at").
		From("users")

	queryWithFilters := utils.ApplyFilters(baseQuery, filters, false)                                            // apply filters
	queryWithFilters = queryWithFilters.Limit(uint64(pagination.Limit)).Offset(uint64(pagination.Offset)) // add pagination

	// generate sql and arguments
	sql, args, err := queryWithFilters.ToSql()
	if err != nil {
		return nil, 0, err
	}

	// rebind to postgresql syntax
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)

	g.logger.Debug("sql query generated", zap.String("query", sql), zap.Any("args", args))

	err = g.DB.SelectContext(ctx, &users, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	// get total users
	var totalUsers int
	countQuery := sq.Select("COUNT(*)").From("users")
	countQueryWithFilters := utils.ApplyFilters(countQuery, filters, true)
	countSQL, countArgs, err := countQueryWithFilters.ToSql()
	if err != nil {
		return nil, 0, err
	}

	// rebind to postgresql syntax
	countSQL = sqlx.Rebind(sqlx.DOLLAR, countSQL)

	g.logger.Debug("sql query generated", zap.String("query", countSQL), zap.Any("args", countArgs))

	err = g.DB.GetContext(ctx, &totalUsers, countSQL, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	return users, totalUsers, nil
}
