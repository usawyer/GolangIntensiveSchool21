package service

import (
	"math"
	"sync"

	"go.uber.org/zap"
)

type ClientStats struct {
	mu    sync.Mutex
	count int
	mean  float64
	m2    float64
	state StateType
}

type StateType int

const (
	StateTypeCollection StateType = iota
	StateTypeDetection
)

func NewClientStats() *ClientStats {
	return &ClientStats{state: StateTypeCollection}
}

func (data *ClientStats) UpdateData(value float64, setSize int) bool {
	data.mu.Lock()
	defer data.mu.Unlock()

	data.count++
	delta := value - data.mean
	data.mean += delta / float64(data.count)
	delta2 := value - data.mean
	data.m2 += delta * delta2

	if data.count == setSize {
		data.state = StateTypeDetection
		return false
	}
	return true
}

func (data *ClientStats) GetState() StateType {
	return data.state
}

func (data *ClientStats) GetMean() float64 {
	data.mu.Lock()
	defer data.mu.Unlock()
	return data.mean
}

func (data *ClientStats) GetSTD() float64 {
	data.mu.Lock()
	defer data.mu.Unlock()

	if data.count < 2 {
		return 0.0
	}

	return math.Sqrt(data.m2 / float64(data.count-1))
}

func (data *ClientStats) LogInfo(logger *zap.Logger) {
	if (data.count+1)%10 == 0 {
		logger.Info("Processed",
			zap.Int("Values", data.count+1),
			zap.Float64("Mean", data.GetMean()),
			zap.Float64("STD", data.GetSTD()))
	}
}
