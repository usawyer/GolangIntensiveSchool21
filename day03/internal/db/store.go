package db

import "search/internal/types"

type Store interface {
	GetPlaces(limit int, offset int) ([]types.Place, int, error)
	GetClosestPlaces(lat float64, lon float64) ([]types.Place, error)
}
