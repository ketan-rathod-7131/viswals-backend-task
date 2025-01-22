package usecase

import (
	"fmt"
	"strconv"

	"github.com/viswals/core/dto"
	"go.uber.org/zap"
)

// ParseCSVRecordToUserData takes a predefined CSV record and parses it to a RawUserData structure.
// This function ensures that even if less data available then move forward.
// NOTE: speciifc to our requirement so kept it under usecase layer
func (c *ProducerUsecase) ParseCSVRecordToUserData(record []string) (dto.RawUserData, error) {

	user := dto.RawUserData{}
	var err error

	// Parses id
	if len(record) >= 1 {
		user.Id, err = strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			c.logger.Error("invalid id in record", zap.Error(err))
			return user, fmt.Errorf("invalid id in record: %w", err)
		}
	}

	// Parse first_name
	if len(record) >= 2 {
		user.FirstName = record[1]
	}

	// Parse first_name
	if len(record) >= 3 {
		user.LastName = record[2]
	}

	// Parse email
	if len(record) >= 4 {
		user.Email = record[3]
	}

	// Parse created_at
	if len(record) >= 5 {
		user.CreatedAt, err = parseTimestamp(record[4])
		if err != nil {
			return user, fmt.Errorf("invalid created_at in record: %w", err)
		}
	}

	// Parse deleted_at
	if len(record) >= 6 {
		user.DeletedAt, err = parseTimestamp(record[5])
		if err != nil {
			return user, fmt.Errorf("invalid deleted_at in record: %w", err)
		}
	}

	// Parse merged_at
	if len(record) >= 7 {
		user.MergedAt, err = parseTimestamp(record[5])
		if err != nil {
			return user, fmt.Errorf("invalid deleted_at in record: %w", err)
		}
	}

	// Parse parent_user_id
	if len(record) >= 8 {
		if record[7] != "" {
			parentUserId, err := strconv.Atoi(record[7])
			if err != nil {
				return user, fmt.Errorf("invalid parent_user_id in record: %w", err)
			}

			user.ParentUserId = int64(parentUserId)
		}
	}

	return user, nil
}

func parseTimestamp(value string) (int64, error) {
	if value == "" {
		return 0, nil // Return 0 if the timestamp is empty
	}

	return strconv.ParseInt(value, 10, 64)
}
