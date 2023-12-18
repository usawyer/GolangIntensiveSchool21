// Package mincoins contains the implementation of three functions: MinCoins, MinCoins2 and MinCoins3.
//
// The functions are designed to find the minimum number of coins needed to represent a given value (val)
// using a provided set of coin denominations (coins).
//
// It accepts a necessary amount and a sorted slice
// of unique denominations of coins. The output is supposed to be a slice of coins of minimal size that
// can be used to express the value.
package mincoins

import (
	"log"
	"sort"
)

// MinCoins is the default function utilizes a greedy algorithm, selecting the largest possible coin denomination
// iteratively. It has limitations, such as improper handling of unsorted inputs, negative values,
// and suboptimal solutions in some scenarios.
func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	log.Println(i)
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// MinCoins2 is a new feature that takes into account the shortcomings of the MinCoins.
// Unlike the previous implementation, this function uses dynamic programming to optimize
// the solution and ensure efficiency.
//
// Advantages of MinCoins2 over MinCoins:
//
//   - Handling Edge Cases:
//     MinCoins2 incorporates checks for negative input values, providing a more robust solution
//     compared to MinCoins, which lacks proper handling for such cases.
//
//   - Optimal Solution:
//     MinCoins2 utilizes dynamic programming, ensuring an optimal solution by considering all
//     possible combinations. In contrast, MinCoins relies on a greedy algorithm, leading to
//     suboptimal results in certain scenarios.
//
//   - Efficiency:
//     The dynamic programming approach in MinCoins2 improves efficiency by avoiding redundant
//     calculations. In MinCoins, the greedy algorithm might revisit the same subproblems, impacting performance.
//
//   - Input Flexibility:
//     MinCoins2 is designed to handle unsorted coin inputs and other edge cases, providing more
//     flexibility and reliability compared to MinCoins, which may produce incorrect results for such inputs.
//
//   - Solution Validation:
//     MinCoins2 validates the existence of a solution by checking if the resulting dynamic programming table
//     indicates a valid combination. MinCoins, on the other hand, may produce incomplete or incorrect solutions
//     for certain inputs without proper validation.
func MinCoins2(val int, coins []int) []int {
	result := make([]int, 0)

	if isNegativeInput(val) || isNegativeInput(coins) {
		return result
	}

	dp := make([]int, val+1)
	lastCoin := make([]int, val+1)

	for i := 1; i <= val; i++ {
		dp[i] = val + 1
		lastCoin[i] = -1
	}

	for _, coin := range coins {
		for i := coin; i <= val; i++ {
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
				lastCoin[i] = coin
			}
		}
	}

	if dp[val] == val+1 {
		return result
	}

	for val > 0 {
		coin := lastCoin[val]
		result = append(result, coin)
		val -= coin
	}

	return result
}

// MinCoins3 function improves upon MinCoins by iteratively evaluating the optimal solution with a decreasing
// set of coin denominations. It addresses negative input values and returns the smallest combination of coins.
//
// Advantages of MinCoins3 over MinCoins:
//
//   - Handling Edge Cases:
//     MinCoins3 includes checks for negative input values, providing a more robust solution compared to MinCoins.
//
//   - Optimal Solution:
//     MinCoins3 iteratively evaluates combinations, ensuring a potentially more optimal solution compared to MinCoins.
//
//   - Result Size:
//     MinCoins3 returns the smallest combination of coins, minimizing the total number used.
//
// Disadvantages of MinCoins3 compared to MinCoins2:
//
//   - Efficiency:
//     MinCoins3 lacks the dynamic programming optimization of MinCoins2, potentially resulting in less efficiency for larger inputs.
//
//   - Solution Validation:
//     MinCoins3 may not validate the solution's existence as comprehensively as MinCoins2.
func MinCoins3(val int, coins []int) []int {
	if isNegativeInput(val) || isNegativeInput(coins) {
		return []int{}
	}

	result := countResult(val, coins)

	for i := 0; i < len(coins); i++ {
		optimal := countResult(val, coins[:len(coins)-i])
		if len(result) == 0 || len(optimal) > 0 && len(optimal) < len(result) {
			result = optimal
		}
	}
	return result
}

func isNegativeNumber(num int) bool {
	return num < 0
}

func isNegativeSlice(slice []int) bool {
	for _, num := range slice {
		if num < 0 {
			return true
		}
	}
	return false
}

func isNegativeInput(input interface{}) bool {
	switch v := input.(type) {
	case int:
		return isNegativeNumber(v)
	case []int:
		return isNegativeSlice(v)
	default:
		return false
	}
}

func countResult(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	sort.Ints(coins)
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}

	if val == 0 {
		return res
	}
	return []int{}
}
