package ex03

import (
	hp "day05/internal/heap"
)

func GrabPresents(presents []hp.Present, capacity int) []hp.Present {
	dp := make([][]int, len(presents)+1)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(presents); i++ {
		for j := 1; j <= capacity; j++ {
			if presents[i-1].Size <= j {
				dp[i][j] = maximum(presents[i-1].Value+dp[i-1][j-presents[i-1].Size], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	result := make([]hp.Present, 0)
	i, j := len(presents), capacity
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			result = append(result, presents[i-1])
			j -= presents[i-1].Size
		}
		i--
	}

	return result
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
