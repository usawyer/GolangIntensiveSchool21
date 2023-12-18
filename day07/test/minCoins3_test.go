package test

import (
	"day07/internal"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestMinCoins3Easy(t *testing.T) {
	coins := []int{1, 5, 10}
	actual := mincoins.MinCoins3(13, coins)
	assert.Equal(t, []int{10, 1, 1, 1}, actual)
}

func TestMinCoins3Duplicated(t *testing.T) {
	coins := []int{1, 5, 1, 10}
	actual := mincoins.MinCoins3(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{1, 5}, actual)
}

func TestMinCoins3Unsorted(t *testing.T) {
	coins := []int{1, 5, 3}
	actual := mincoins.MinCoins3(10, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{5, 5}, actual)
}

func TestMinCoins3Zero(t *testing.T) {
	coins := []int{1, 5, 10}
	actual := mincoins.MinCoins3(0, coins)
	assert.Equal(t, []int{}, actual)
}

func TestMinCoins3Empty(t *testing.T) {
	actual := mincoins.MinCoins3(13, []int{})
	assert.Equal(t, []int{}, actual)
}

func TestMinCoins3Negative(t *testing.T) {
	coins := []int{1, -2, 3}
	timeout := time.After(3 * time.Second)
	done := make(chan []int)
	go func() {
		actual := mincoins.MinCoins3(5, coins)
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

func TestMinCoins3NegativeVal(t *testing.T) {
	coins := []int{1, 2, 3}
	timeout := time.After(3 * time.Second)
	done := make(chan []int)
	go func() {
		actual := mincoins.MinCoins3(-5, coins)
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

func TestMinCoins3Impossible(t *testing.T) {
	coins := []int{4, 6, 10}
	actual := mincoins.MinCoins3(2, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{}, actual)

	actual = mincoins.MinCoins3(11, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{}, actual)
}

func TestMinCoins3Optimal(t *testing.T) {
	coins := []int{1, 2, 3, 10}
	actual := mincoins.MinCoins3(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{3, 3}, actual)

	coins = []int{4, 1, 3, 7}
	actual = mincoins.MinCoins3(6, coins)
	sort.Ints(actual)
	assert.Equal(t, []int{3, 3}, actual)
}
