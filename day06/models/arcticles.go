package models

import "time"

type Articles struct {
	Id      int       `json:"id" gorm:"primaryKey"`
	Tittle  string    `json:"tittle"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
