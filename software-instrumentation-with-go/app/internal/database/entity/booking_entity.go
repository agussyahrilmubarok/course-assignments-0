package entity

import (
	"time"
)

type Booking struct {
	ID           string `gorm:"primaryKey"`
	BatchID      string `gorm:"index;not null"`
	Code         string
	Status       int
	ReservedAt   time.Time
	ExpiredAt    time.Time
	PaidAt       time.Time
	FailedAt     time.Time
	CustomerName string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Batch Batch `gorm:"foreignKey:BatchID;references:ID"`
}
