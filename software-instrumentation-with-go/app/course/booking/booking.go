package booking

import (
	"app/course/catalog"
	"app/internal/database/entity"
	"time"
)

type Booking struct {
	ID           string
	Code         string
	Status       Status
	ReservedAt   time.Time
	ExpiredAt    time.Time
	PaidAt       time.Time
	FailedAt     time.Time
	CustomerName string
	Batch        *catalog.Batch
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Status int

const (
	StatusUnknown Status = iota
	StatusCreated
	StatusReserved
	StatusCompleted
	StatusFailed
	StatusExpired
)

func ToBooking(e *entity.Booking) *Booking {
	if e == nil {
		return nil
	}

	b := &Booking{
		ID:           e.ID,
		Code:         e.Code,
		Status:       Status(e.Status),
		ReservedAt:   e.ReservedAt,
		ExpiredAt:    e.ExpiredAt,
		PaidAt:       e.PaidAt,
		FailedAt:     e.FailedAt,
		CustomerName: e.CustomerName,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}

	if e.Batch.ID != "" {
		b.Batch = catalog.ToBatch(&e.Batch)
	}

	return b
}

func ToBookingEntity(b *Booking) *entity.Booking {
	if b == nil {
		return nil
	}

	var batchID string

	if b.Batch != nil {
		batchID = b.Batch.ID
	}

	return &entity.Booking{
		ID:           b.ID,
		Code:         b.Code,
		Status:       int(b.Status),
		BatchID:      batchID,
		ReservedAt:   b.ReservedAt,
		ExpiredAt:    b.ExpiredAt,
		PaidAt:       b.PaidAt,
		FailedAt:     b.FailedAt,
		CustomerName: b.CustomerName,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}
