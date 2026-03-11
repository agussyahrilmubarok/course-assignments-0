package entity

import "time"

type Batch struct {
	ID             string `gorm:"primaryKey"`
	CourseID       string `gorm:"index;not null"`
	Code           string
	Name           string
	PriceAmount    int64
	PriceCurrency  string
	MaxSeats       int32
	AvailableSeats int32
	Status         int
	StartDate      int64
	EndDate        int64
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Course Course `gorm:"foreignKey:CourseID;references:ID"`
}
