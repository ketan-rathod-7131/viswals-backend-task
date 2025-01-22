package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_database "github.com/viswals/consumer/usecase/repository/database/mock"
	"github.com/viswals/core/models"
)

// TODO: Write test cases for GetAllUsers, GetUserById and other crud APIs.

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConsumerRepo := mock_database.NewMockIConsumerRepository(ctrl)
	defer ctrl.Finish()

	ctx := context.Background()

	tests := []struct {
		name     string
		user     models.User
		wantID   string
		wantErr  bool
		mockFunc func()
	}{
		{
			name: "success",
			user: models.User{
				Email:     "test@example.com",
				FirstName: "Test",
				LastName:  "User",
			},
			wantID:  "123",
			wantErr: false,
			mockFunc: func() {
				mockConsumerRepo.EXPECT().
					CreateUser(ctx, gomock.Any()).
					Return("123", nil)
			},
		},
		{
			name: "duplicate Email",
			user: models.User{
				Email:     "existing@example.com",
				FirstName: "Test",
				LastName:  "User",
			},
			wantID:  "",
			wantErr: true,
			mockFunc: func() {
				mockConsumerRepo.EXPECT().
					CreateUser(ctx, gomock.Any()).
					Return("", fmt.Errorf("duplicate email id"))
			},
		},
		{
			name: "invalid Email",
			user: models.User{
				Email:     "invalid-email",
				FirstName: "Test",
				LastName:  "User",
			},
			wantID:  "",
			wantErr: true,
			mockFunc: func() {
				mockConsumerRepo.EXPECT().
					CreateUser(ctx, gomock.Any()).
					Return("", fmt.Errorf("invalid email format"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// call mockFunc to setup mocks and call createUser function for testing
			tt.mockFunc()
			id, err := mockConsumerRepo.CreateUser(ctx, tt.user)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantID, id)
		})
	}
}
