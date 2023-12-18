package internal

import (
	"log"
	"sort"
)

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
