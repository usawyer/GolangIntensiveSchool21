package test

import (
	hp "day05/pkg/heap"
	"reflect"
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

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", actual, expected)
	}
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

	if err == nil {
		t.Error("Expected an error but got none.")
		return
	}

	if err.Error() != expectedError {
		t.Errorf("Error message was incorrect, got: %v, want: %v.", err.Error(), expectedError)
		return
	}

	if len(actual) > 0 {
		t.Errorf("Expected an empty result, but got: %v.", actual)
		return
	}
}
