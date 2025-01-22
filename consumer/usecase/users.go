package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/viswals/core/infrastructure/redis"
	"github.com/viswals/core/models"
	"github.com/viswals/core/pkg/utils"
	"go.uber.org/zap"
)

func (c *ConsumerUsecase) GetAllUsers(ctx context.Context, pagination utils.PaginationParams, filters []utils.Filter) (users []models.User, totalUsers int, err error) {

	users, totalUsers, err = c.db.GetAllUsers(ctx, pagination, filters)
	if err != nil {
		c.logger.Error("Failed to get user data", zap.Error(err))
		return users, totalUsers, err
	}

	// decrypt email
	for i, u := range users {
		decryptedEmail, err := c.em.Decrypt(u.Email)
		if err != nil {
			c.logger.Error("Failed to decrypt user email", zap.Error(err))
			users[i].Email = "" // do not expose internal data format of email
		} else {
			users[i].Email = decryptedEmail
		}
	}

	return users, totalUsers, nil
}

func (c *ConsumerUsecase) GetUserById(ctx context.Context, id int64) (user models.User, err error) {

	// fetch data from cache service
	data, err := c.cm.Get(ctx, redis.GetKey("users", id))
	if err != nil {
		// fetch data from repository service
		user, err := c.db.GetUserById(ctx, id)
		if err != nil {
			c.logger.Error("Failed to get user data", zap.Error(err))
			return user, err
		}

		// set data to the cache service
		data, err := json.Marshal(&user)
		if err != nil {
			c.logger.Error("Failed to marshal user data", zap.Error(err))
		}

		err = c.cm.Set(ctx, redis.GetKey("users", id), string(data), time.Minute*10)
		if err != nil {
			c.logger.Error("Failed to set cache", zap.Error(err))
		}

		// decrypt email from the database
		decryptedEmail, err := c.em.Decrypt(user.Email)
		if err != nil {
			c.logger.Error("Failed to decrypt user email", zap.Error(err))
			user.Email = "" // do not expose ineternal data format of email
			return user, err
		}
		user.Email = decryptedEmail

		return user, nil
	}

	// cache hit
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		c.logger.Error("Failed to unmarshal user data", zap.Error(err))
		c.cm.Delete(ctx, redis.GetKey("users", id)) // remove the invalid cache entry
		return user, err
	}

	// decrypt the user email
	decryptedEmail, err := c.em.Decrypt(user.Email)
	if err != nil {
		c.logger.Error("Failed to decrypt user email", zap.Error(err))
		user.Email = "" // do not expose inetrnal data format of email
		return user, err
	}

	user.Email = decryptedEmail

	return user, nil
}
