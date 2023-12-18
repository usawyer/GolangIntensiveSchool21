package test

import (
	"day07/internal"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestMinCoinsEasy(t *testing.T) {
	coins := []int{1, 5, 10}
	actual := internal.MinCoins(13, coins)
	assert.Equal(t, []int{10, 1, 1, 1}, actual)
}

func TestMinCoinsDuplicated(t *testing.T) {
	coins := []int{1, 5, 1, 10}
	actual := internal.MinCoins(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{1, 5}, actual)
}

func TestMinCoinsUnsorted(t *testing.T) {
	coins := []int{1, 5, 3}
	actual := internal.MinCoins(10, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{5, 5}, actual)
}

func TestMinCoinsZero(t *testing.T) {
	coins := []int{1, 5, 10}
	actual := internal.MinCoins(0, coins)
	assert.Equal(t, []int{}, actual)
}

func TestMinCoinsEmpty(t *testing.T) {
	actual := internal.MinCoins(13, []int{})
	assert.Equal(t, []int{}, actual)
}

func TestMinCoinsNegative(t *testing.T) {
	coins := []int{1, -2, 3}
	timeout := time.After(3 * time.Second)
	done := make(chan []int)
	go func() {
		actual := internal.MinCoins(5, coins)
		done <- actual
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case actual := <-done:
		assert.Equal(t, []int{}, actual)
	}
}

func TestMinCoinsNegativeVal(t *testing.T) {
	coins := []int{1, 2, 3}
	timeout := time.After(3 * time.Second)
	done := make(chan []int)
	go func() {
		actual := internal.MinCoins(-5, coins)
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

func TestMinCoinsImpossible(t *testing.T) {
	coins := []int{4, 6, 10}
	actual := internal.MinCoins(2, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{}, actual)

	actual = internal.MinCoins(11, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{}, actual)
}

func TestMinCoinsOptimal(t *testing.T) {
	coins := []int{1, 2, 3, 10}
	actual := internal.MinCoins(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{3, 3}, actual)

	coins = []int{4, 1, 3, 7}
	actual = internal.MinCoins(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{3, 3}, actual)
}
