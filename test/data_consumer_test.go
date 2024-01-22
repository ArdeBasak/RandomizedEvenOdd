package test

import (
	"EvenOdd/mocks"
	"EvenOdd/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestConsumeDataForSingleInput tests the ConsumeData function with a single input.
func TestConsumeDataForSingleInput(t *testing.T) {
	// Create an instance of the mock
	mock := &mocks.DataProvider{}

	// Set the expected return value for the GetRandomNumber method
	mock.On("GetRandomNumber", 5).Return(3, nil).Once()

	// Your code that uses the DataProvider interface
	result, err := service.ConsumeData(mock, 5)

	// Assert that the result and error are as expected
	assert.Equal(t, "Odd", result)
	assert.NoError(t, err)

	// Assert that the GetRandomNumber method was called with the expected input
	mock.AssertExpectations(t)
}

// TestConsumeDataForMultipleInput tests the ConsumeData function with multiple values.
func TestConsumeDataForMultipleInput(t *testing.T) {
	// Create an instance of the mock
	mock := &mocks.DataProvider{}

	// Set the expected return values for the GetRandomNumber method
	mock.On("GetRandomNumber", 20).Return(10, nil).Once()
	mock.On("GetRandomNumber", 30).Return(15, nil).Once()
	mock.On("GetRandomNumber", 10).Return(5, nil).Once()

	// Set the expected return value for an error scenario
	mock.On("GetRandomNumber", -1).Return(0, errors.New("Invalid id")).Once()

	// Your code that uses the DataProvider interface
	testCases := []struct {
		input int
		want  string
		err   error
	}{
		{input: 20, want: "Even", err: nil},
		{input: 30, want: "Odd", err: nil},
		{input: 10, want: "Odd", err: nil},
		{input: -1, want: "", err: errors.New("Invalid id")},
	}

	for _, tc := range testCases {
		result, err := service.ConsumeData(mock, tc.input)

		// Assert that the result and error are as expected
		assert.Equal(t, tc.want, result)
		assert.Equal(t, tc.err, err)
	}

	// Assert that the GetRandomNumber methods were called with the expected inputs
	mock.AssertExpectations(t)
}
