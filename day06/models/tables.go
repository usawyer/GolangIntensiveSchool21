package models

type Articles struct {
	ID      uint `gorm:"primaryKey;column:id;autoIncrement"`
	Tittle  string
	Content string
}
