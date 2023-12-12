package test

import (
	"day05/cmd/ex03"
	hp "day05/internal/heap"
	"testing"
)

func TestKnapsackEasy(t *testing.T) {
	items := []hp.Present{
		{Value: 5, Size: 3},
		{Value: 3, Size: 2},
		{Value: 4, Size: 1},
	}

	indices := ex03.GrabPresents(items, 5)
	value := 0
	for _, i := range indices {
		value += i.Value
	}

	if value != 9 {
		t.Errorf("Expected %d, got %d", 9, value)
	}
}

func TestKnapsackZeroCapacity(t *testing.T) {
	items := []hp.Present{
		{Value: 5, Size: 3},
		{Value: 3, Size: 2},
		{Value: 4, Size: 1},
	}

	indices := ex03.GrabPresents(items, 0)
	value := 0
	for _, i := range indices {
		value += i.Value
	}

	if value != 0 {
		t.Errorf("Expected %d, got %d", 9, value)
	}
}

func TestKnapsackNoItem(t *testing.T) {
	items := []hp.Present{}

	indices := ex03.GrabPresents(items, 5)
	value := 0
	for _, i := range indices {
		value += i.Value
	}

	if value != 0 {
		t.Errorf("Expected %d, got %d", 9, value)
	}
}
