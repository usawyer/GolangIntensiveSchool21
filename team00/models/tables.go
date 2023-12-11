package models

import (
	"time"
)

type Session struct {
	ID         uint   `gorm:"primaryKey;column:id;autoIncrement"`
	SessionId  string `gorm:"column:session_id"`
	ServMean   float64
	ServStd    float64
	ClientMean float64 `gorm:"default:null"`
	ClientStd  float64 `gorm:"default:null"`
	ClientK    float64 `gorm:"default:null"`
	Status     string
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

type Anomaly struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Session   Session `gorm:"foreignKey:ID"`
	SessionID uint    `gorm:"column:session_id"`
	Frequency float64
	Timestamp time.Time `gorm:"type:timestamp;column:timestamp"`
}
