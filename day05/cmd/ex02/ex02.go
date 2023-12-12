package ex02

import (
	"container/heap"
	hp "day05/internal/heap"

	"github.com/pkg/errors"
)

func GetNCoolestPresents(presents []hp.Present, n int) ([]hp.Present, error) {
	if n < 0 || n > len(presents) {
		return nil, errors.New("invalid number of n")
	}

	h := &hp.PresentHeap{}
	heap.Init(h)

	for _, present := range presents {
		heap.Push(h, present)
	}

	result := make([]hp.Present, n)
	for i := 0; i < n; i++ {
		result[i] = heap.Pop(h).(hp.Present)
	}

	return result, nil
}
