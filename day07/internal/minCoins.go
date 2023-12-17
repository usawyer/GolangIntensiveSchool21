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
	return res
}
