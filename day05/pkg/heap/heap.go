package myHeap

import (
	"container/heap"
	"github.com/pkg/errors"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h *PresentHeap) GetNCoolestPresents(n int) ([]Present, error) {
	if n < 0 || n > len(*h) {
		return nil, errors.New("invalid number of n")
	}

	presentHeap := &PresentHeap{}
	heap.Init(presentHeap)

	for _, present := range *h {
		heap.Push(presentHeap, present)
	}

	result := make([]Present, n)
	for i := 0; i < n; i++ {
		result[i] = heap.Pop(presentHeap).(Present)
	}

	return result, nil
}

func (h *PresentHeap) GrabPresents(capacity int) []Present {
	dp := make([][]int, len(*h)+1)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(*h); i++ {
		for j := 1; j <= capacity; j++ {
			if (*h)[i-1].Size <= j {
				dp[i][j] = maximum((*h)[i-1].Value+dp[i-1][j-(*h)[i-1].Size], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	result := make([]Present, 0)
	i, j := len(*h), capacity
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			result = append(result, (*h)[i-1])
			j -= (*h)[i-1].Size
		}
		i--
	}

	return result
}

func (h *PresentHeap) Len() int {
	return len(*h)
}

func (h *PresentHeap) Less(i, j int) bool {
	if (*h)[i].Value == (*h)[j].Value {
		return (*h)[i].Size < (*h)[j].Size
	}
	return (*h)[i].Value > (*h)[j].Value
}

func (h *PresentHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *PresentHeap) Push(x interface{}) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
