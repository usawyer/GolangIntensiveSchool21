package test

import (
	"day07/internal"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestMinCoins2Easy(t *testing.T) {
	coins := []int{1, 5, 10}
	actual := internal.MinCoins2(13, coins)
	assert.Equal(t, []int{10, 1, 1, 1}, actual)
}

func TestMinCoins2Duplicated(t *testing.T) {
	coins := []int{1, 5, 1, 10}
	actual := internal.MinCoins2(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{1, 5}, actual)
}

func TestMinCoins2Unsorted(t *testing.T) {
	coins := []int{1, 5, 3}
	actual := internal.MinCoins2(10, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{5, 5}, actual)
}

func TestMinCoins2Zero(t *testing.T) {
	coins := []int{1, 5, 10}
	actual := internal.MinCoins2(0, coins)
	assert.Equal(t, []int{}, actual)
}

func TestMinCoins2Empty(t *testing.T) {
	actual := internal.MinCoins2(13, []int{})
	assert.Equal(t, []int{}, actual)
}

func TestMinCoins2Negative(t *testing.T) {
	coins := []int{1, -2, 3}
	timeout := time.After(3 * time.Second)
	done := make(chan []int)
	go func() {
		actual := internal.MinCoins2(5, coins)
		done <- actual
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case actual := <-done:
		sort.Ints(actual)
		assert.Equal(t, []int{1, 1, 3}, actual)
	}
}

func TestMinCoins2NegativeVal(t *testing.T) {
	coins := []int{1, 2, 3}
	timeout := time.After(3 * time.Second)
	done := make(chan []int)
	go func() {
		actual := internal.MinCoins2(-5, coins)
		done <- actual
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case actual := <-done:
		sort.Ints(actual)
		assert.Equal(t, []int{}, actual)
	}
}

func TestMinCoins2Impossible(t *testing.T) {
	coins := []int{4, 6, 10}
	actual := internal.MinCoins2(2, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{}, actual)

	actual = internal.MinCoins2(8, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{}, actual)
}

func TestMinCoins2Optimal(t *testing.T) {
	coins := []int{1, 2, 3, 10}
	actual := internal.MinCoins2(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{3, 3}, actual)
}
