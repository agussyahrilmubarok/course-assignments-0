package entity

import "time"

type Course struct {
	ID        string `gorm:"primaryKey"`
	Code      string `gorm:"uniqueIndex;not null"`
	Title     string `gorm:"not null"`
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time

	Batches []Batch `gorm:"foreignKey:CourseID;references:ID"`
}
