package usecase

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/viswals/core/dto"
	"github.com/viswals/core/infrastructure/rabbitmq"
	"go.uber.org/zap"
)

func (c *ProducerUsecase) PublishUserDataToQueue(ctx context.Context, queue string, userdata dto.RawUserData) error {

	message, err := json.Marshal(userdata)
	if err != nil {
		c.logger.Error("failed to marshal user data", zap.Error(err))
		return err
	}

	// publish message to RabbitMQ queue
	err = c.rmq.PublishWithContext(ctx, rabbitmq.WithContentType("application/json"), rabbitmq.WithBody(message), rabbitmq.WithRoutingKey(queue))
	if err != nil {
		c.logger.Error("failed to publish message to queue", zap.Error(err))
		return err
	}

	c.logger.Info("message published to queue", zap.String("user", string(message)))

	return nil
}

func (c *ProducerUsecase) PublishCSVDataToQueue(filepath string, queue string) error {
	file, err := os.Open(filepath)
	if err != nil {
		c.logger.Error("failed to open CSV file", zap.Error(err))
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// buffered channel to hold users data
	userDataCh := make(chan dto.RawUserData, 100)

	// number of worker goroutines
	workerCount := 5

	// waitgroup to track goroutines
	var wg sync.WaitGroup

	// goroutine for reading data from CSV file and writing data back into the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(userDataCh) // close channel when done reading

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.logger.Error("failed to read CSV record", zap.Error(err))
				continue
			}

			userData, err := c.ParseCSVRecordToUserData(record)
			if err != nil {
				c.logger.Error("failed to parse CSV record", zap.Error(err))
				continue
			}

			userDataCh <- userData
		}
	}()

	// worker goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for userData := range userDataCh {
				err := c.PublishUserDataToQueue(context.Background(), queue, userData)
				if err != nil {
					c.logger.Error("failed to publish user data to queue", zap.Error(err))
				} else {
					c.logger.Info("successfully published user data to queue")
				}
			}
		}(i)
	}

	wg.Wait() // wait for all goroutines to finish
	c.logger.Info("All CSV data processed and published to queue")
	return nil
}
