package catalog

import (
	"app/internal/database/entity"
	"time"
)

type Batch struct {
	ID             string
	Code           string
	Name           string
	Price          Money
	MaxSeats       int32
	AvailableSeats int32
	Status         BatchStatus
	StartDate      time.Time
	EndDate        time.Time
	Course         *Course
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BatchStatus int

const (
	BatchStatusDraft BatchStatus = iota
	BatchStatusPublished
	BatchStatusArchived
)

func (b Batch) IsPublished() bool {
	return b.Status == BatchStatusPublished
}

func (b *Batch) Publish() error {
	if b.Status != BatchStatusDraft {
		return ErrInvalidBatchStatus
	}

	b.Status = BatchStatusPublished
	b.UpdatedAt = time.Now()

	return nil
}

func (b *Batch) Archive() error {
	if b.Status == BatchStatusArchived {
		return ErrInvalidBatchStatus
	}

	b.Status = BatchStatusArchived
	b.UpdatedAt = time.Now()

	return nil
}

func (b Batch) IsFull() bool {
	return b.AvailableSeats <= 0
}

func (b *Batch) BookSeat(qty int32) error {
	if b.Status != BatchStatusPublished {
		return ErrBatchNotOpen
	}

	if b.AvailableSeats < qty {
		return ErrNotEnoughSeats
	}

	if time.Now().After(b.StartDate) {
		return ErrBatchAlreadyStarted
	}

	b.AvailableSeats -= qty
	b.UpdatedAt = time.Now()

	return nil
}

func (b *Batch) CancelSeat(qty int32) {
	b.AvailableSeats += qty
	b.UpdatedAt = time.Now()
}

func (b Batch) HasStarted() bool {
	return time.Now().After(b.StartDate)
}

func ToBatch(e *entity.Batch) *Batch {
	if e == nil {
		return nil
	}

	b := &Batch{
		ID:   e.ID,
		Code: e.Code,
		Name: e.Name,
		Price: Money{
			Amount:   e.PriceAmount,
			Currency: e.PriceCurrency,
		},
		MaxSeats:       e.MaxSeats,
		AvailableSeats: e.AvailableSeats,
		Status:         BatchStatus(e.Status),
		StartDate:      time.Unix(e.StartDate, 0),
		EndDate:        time.Unix(e.EndDate, 0),
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}

	if e.Course.ID != "" {
		b.Course = ToCourse(&e.Course)
	}

	return b
}

func ToBatchEntity(b *Batch) *entity.Batch {
	if b == nil {
		return nil
	}

	var courseID string

	if b.Course != nil {
		courseID = b.Course.ID
	}

	return &entity.Batch{
		ID:             b.ID,
		CourseID:       courseID,
		Code:           b.Code,
		Name:           b.Name,
		PriceAmount:    b.Price.Amount,
		PriceCurrency:  b.Price.Currency,
		MaxSeats:       b.MaxSeats,
		AvailableSeats: b.AvailableSeats,
		Status:         int(b.Status),
		StartDate:      b.StartDate.Unix(),
		EndDate:        b.EndDate.Unix(),
		CreatedAt:      b.CreatedAt,
		UpdatedAt:      b.UpdatedAt,
	}
}
