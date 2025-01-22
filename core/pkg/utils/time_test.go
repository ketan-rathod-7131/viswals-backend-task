package utils

import (
	"testing"
	"time"
)

func TestGetTimeFromEpoch(t *testing.T) {
	tests := []struct {
		name      string
		epoch     int64
		expectErr bool
		expected  *time.Time
	}{
		{
			name:      "valid epoch time",
			epoch:     1672531200000, // Corresponds to 2023-01-01T00:00:00Z
			expectErr: false,
			expected: func() *time.Time {
				t := time.Unix(1672531200, 0)
				return &t
			}(),
		},
		{
			name:      "invalid epoch time (negative)",
			epoch:     -1000,
			expectErr: true,
			expected:  nil,
		},
		{
			name:      "invalid epoch time (zero)",
			epoch:     0,
			expectErr: true,
			expected:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := GetTimeFromEpoch(test.epoch)

			if test.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if result == nil {
					t.Errorf("expected a time but got nil")
				} else if !result.Equal(*test.expected) {
					t.Errorf("expected %v but got %v", test.expected, result)
				}
			}
		})
	}
}
