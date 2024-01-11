package benchmark_test

import (
	mincoins "day07/internal"
	"math/rand"
	"testing"
)

var (
	data = randomGenerator(99999)
)

func randomGenerator(size int) []int {
	randNumber := make([]int, size)
	for i := 0; i < size; i++ {
		randNumber[i] = rand.Intn(100)
	}

	return randNumber
}

func BenchmarkMinCoins2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mincoins.MinCoins2(1234, data)
	}
}

//func BenchmarkMinCoins(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		mincoins.MinCoins3(1234, data)
//	}
//}
