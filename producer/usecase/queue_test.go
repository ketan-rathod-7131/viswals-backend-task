package usecase_test

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_interfaces "github.com/viswals/core/interfaces/mocks"
	"github.com/viswals/producer/usecase"
)

// TODO: prepare tabuler test cases for testing producer services

func TestPublishCSVDataToQueue_Success_ValidCSVFile(t *testing.T) {

	csvContent := `8,Hanah,Schmidt,Hanah_Schmidt1965@gmail.edu,1361218223000,-1,-1,-1
31,Emily,Tamm,EmilyTamm@gmail.edu,1361367320000,-1,-1,-1`

	tmpfile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(csvContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// setup mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQueue := mock_interfaces.NewMockIQueueService(ctrl)
	mockLogger := mock_interfaces.NewMockILogger(ctrl)

	// set up expectations for the mock logger and queue service
	mockLogger.EXPECT().Info("message published to queue", gomock.Any()).Times(2)
	// for each data row, we expect PublishWithContext to be called twice
	mockQueue.EXPECT().PublishWithContext(gomock.Any(), gomock.Any()).Return(nil).Times(2)
	// expect success logs for the published messages
	mockLogger.EXPECT().Info(gomock.Any()).Times(2)
	// expect the final completion message
	mockLogger.EXPECT().Info("All CSV data processed and published to queue").Times(1)

	// initialize the producer with the mocks
	producer := usecase.New(mockQueue, usecase.WithLogger(mockLogger))

	// run the test: publish CSV data to the queue
	err = producer.PublishCSVDataToQueue(tmpfile.Name(), "test_queue")

	// check if the error is as expected (no error for success case)
	assert.NoError(t, err)
}

func TestPublishCSVDataToQueue_Failure_InValidCSVFile(t *testing.T) {

	// prepare csv with invalid content
	csvContent := `8,Hanah,Schmidt,Hanah_Schmidt1965@gmail.edu,1361218223000,-1,-1,-1
invalid-id,Emily,Tamm,EmilyTamm@gmail.edu,1361367320000,-1,-1,-1`

	// temp file for testing
	tmpfile, err := os.CreateTemp("", "test-*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// write the CSV content to the file
	if _, err := tmpfile.Write([]byte(csvContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// setup mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQueue := mock_interfaces.NewMockIQueueService(ctrl)
	mockLogger := mock_interfaces.NewMockILogger(ctrl)

	mockLogger.EXPECT().Error(gomock.Any(), gomock.Any()).Times(2)
	mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(2)
	// we expect this to be called once, as out of two records one is buggy.
	mockQueue.EXPECT().PublishWithContext(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	// expect the final completion message
	mockLogger.EXPECT().Info(gomock.Any()).Times(1)

	// initialize the producer with the mocks
	producer := usecase.New(mockQueue, usecase.WithLogger(mockLogger))

	// run the test: publish CSV data to the queue
	err = producer.PublishCSVDataToQueue(tmpfile.Name(), "test_queue")

	// there should be no error, as it logs the errors if any.
	assert.NoError(t, err)
}
