package usecase

import (
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	coreDto "github.com/viswals/core/dto"
	"github.com/viswals/core/infrastructure/rabbitmq"
	"github.com/viswals/core/models"
	"github.com/viswals/core/pkg/utils"
	"go.uber.org/zap"
)

func (c *ConsumerUsecase) ConsumeUserData(ctx context.Context, queue string) {

	err := c.rmq.Qos(10, 0, false) // Prefetch up to 10 messages per consumer
	if err != nil {
		c.logger.Error("failed to set QoS for consumer", zap.Error(err))
		return
	}

	consumeCh, err := c.rmq.ConsumeWithContext(ctx, queue, rabbitmq.WithAutoAck(false))
	if err != nil {
		c.logger.Error("failed to consume messages from queue", zap.Error(err))
		return
	}

	// Worker pool configuration
	workerCount := 10
	messageChan := make(chan amqp091.Delivery, workerCount)

	// Start worker goroutines
	for i := 0; i < workerCount; i++ {
		go c.worker(ctx, messageChan)
	}

	// Feed messages into the worker pool
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("Context done, stopping consumer")
			close(messageChan) // Stop workers gracefully
			return
		case msg, ok := <-consumeCh:
			if !ok {
				c.logger.Warn("Consumer channel closed, stopping")
				close(messageChan) // Stop workers gracefully
				return
			}

			messageChan <- msg
		}
	}
}

// worker processes messages concurrently
func (c *ConsumerUsecase) worker(ctx context.Context, messageChan <-chan amqp091.Delivery) {
	for msg := range messageChan {
		rawUserData, err := c.GetRawUserDataFromMessage(ctx, msg.Body)
		if err != nil {
			c.logger.Error("failed to process message", zap.Error(err))
			msg.Nack(false, true)
			continue
		}

		// Convert raw user data into database model
		user, err := c.ParseRawUserData(ctx, rawUserData)
		if err != nil {
			c.logger.Error("failed to parse raw user data", zap.Error(err))
			msg.Nack(false, true)
			continue
		}

		// Encrypt data before storing it in the database
		encryptedUserEmail, err := c.em.Encrypt(user.Email)
		if err != nil {
			c.logger.Error("failed to encrypt user email", zap.String("email", user.Email), zap.Error(err))
			msg.Nack(false, true)
			continue
		}

		// Update user email to encrypted user email
		user.Email = encryptedUserEmail

		userId, err := c.db.CreateUser(ctx, user)
		if err != nil {
			c.logger.Error("failed to create user in database", zap.Error(err))
			msg.Nack(false, true)
			continue
		}

		// Acknowledge the message after successful processing
		err = msg.Ack(false)
		if err != nil {
			c.logger.Error("failed to acknowledge message", zap.Error(err))
		}

		c.logger.Info("Processed and saved user data", zap.String("user_id", userId))
	}
}

// GetRawUserDataFromMessage processes the incoming message body (implement your logic here)
func (c *ConsumerUsecase) GetRawUserDataFromMessage(ctx context.Context, messageBody []byte) (userData coreDto.RawUserData, err error) {

	err = json.Unmarshal(messageBody, &userData)
	if err != nil {
		return userData, err
	}

	return userData, nil
}

func (p *ConsumerUsecase) ParseRawUserData(ctx context.Context, rawUserData coreDto.RawUserData) (user models.User, err error) {
	user.Email = rawUserData.Email
	user.FirstName = rawUserData.FirstName
	user.LastName = rawUserData.LastName

	if rawUserData.ParentUserId >= 0 {
		user.ParentUserId = &rawUserData.ParentUserId
	}

	createdAt, err := utils.GetTimeFromEpoch(rawUserData.CreatedAt)
	if err == nil {
		user.CreatedAt = createdAt
	}

	deletedAt, err := utils.GetTimeFromEpoch(rawUserData.DeletedAt)
	if err == nil {
		user.DeletedAt = deletedAt
	}

	mergedAt, err := utils.GetTimeFromEpoch(rawUserData.MergedAt)
	if err == nil {
		user.MergedAt = mergedAt
	}

	return user, nil
}
