package test

import (
	hp "day05/pkg/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapEasy(t *testing.T) {
	unsortedPresents := hp.PresentHeap{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	expected := []hp.Present{{Value: 5, Size: 1}, {Value: 5, Size: 2}}
	actual, _ := unsortedPresents.GetNCoolestPresents(2)
	assert.Equal(t, expected, actual)
}

func TestHeapIncorrectN(t *testing.T) {
	unsortedPresents := hp.PresentHeap{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	expectedError := "invalid number of n"
	actual, err := unsortedPresents.GetNCoolestPresents(6)

	assert.Error(t, err, "Expected an error but got none.")
	assert.EqualError(t, err, expectedError, "Error message was incorrect.")
	assert.Empty(t, actual, "Expected an empty result, but got: %v.", actual)
}
