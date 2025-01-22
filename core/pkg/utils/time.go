package utils

import (
	"fmt"
	"time"
)

func GetTimeFromEpoch(epoch int64) (*time.Time, error) {

	if epoch <= 0 {
		return nil, fmt.Errorf("invalid epoch time: %d", epoch)
	}

	// Convert epoch time (milliseconds) to seconds
	epochInSeconds := epoch / 1000
	t := time.Unix(epochInSeconds, 0)

	return &t, nil
}
