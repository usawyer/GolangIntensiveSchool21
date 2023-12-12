package test

import (
	hp "day05/pkg/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsackEasy(t *testing.T) {
	items := hp.PresentHeap{
		{Value: 5, Size: 3},
		{Value: 3, Size: 2},
		{Value: 4, Size: 1},
	}

	expected := []hp.Present{{Value: 4, Size: 1}, {Value: 5, Size: 3}}
	assert.Equal(t, expected, items.GrabPresents(5))
}

func TestKnapsackZeroCapacity(t *testing.T) {
	items := hp.PresentHeap{
		{Value: 5, Size: 3},
		{Value: 3, Size: 2},
		{Value: 4, Size: 1},
	}

	expected := []hp.Present{}
	assert.Equal(t, expected, items.GrabPresents(0))
}

func TestKnapsackNoItem(t *testing.T) {
	items := hp.PresentHeap{}

	expected := []hp.Present{}
	assert.Equal(t, expected, items.GrabPresents(0))
}
