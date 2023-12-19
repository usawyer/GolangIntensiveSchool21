package test

import (
	"day08/internal/arithmetic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetElementValid(t *testing.T) {
	arr := []int{1, 10, 100, 10000, 100000, -1}

	for i := 0; i < len(arr); i++ {
		res, err := arithmetic.GetElement(arr, i)
		assert.NoError(t, err)
		assert.Equal(t, arr[i], res)
	}
}

func TestGetElementInvalid(t *testing.T) {
	t.Run("Case№1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		_, err := arithmetic.GetElement(arr, -2)
		assert.Error(t, err)
	})

	t.Run("Case№2", func(t *testing.T) {
		arr := []int{1, 10, 100, 10000, 100000}
		_, err := arithmetic.GetElement(arr, len(arr))
		assert.Error(t, err)
	})

	t.Run("Case№3", func(t *testing.T) {
		arr := []int{1, 10, 100, 10000, 100000}
		_, err := arithmetic.GetElement(arr, len(arr)+500)
		assert.Error(t, err)
	})

	t.Run("Case№4", func(t *testing.T) {
		_, err := arithmetic.GetElement(nil, 2)
		assert.Error(t, err)
	})
}
