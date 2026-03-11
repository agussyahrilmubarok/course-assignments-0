package booking

import (
	"app/course/catalog"
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	bookingStore *Store
	catalogStore *catalog.Store
}

func NewService(
	bookingStore *Store,
	catalogStore *catalog.Store,
) *Service {
	return &Service{
		bookingStore: bookingStore,
		catalogStore: catalogStore,
	}
}

func (s *Service) CreateBooking(ctx context.Context, batchCode string, customerName string) (*Booking, error) {
	if customerName == "" {
		return nil, ErrInvalidCustomerName
	}

	batch, err := s.catalogStore.FindBatchCourseByBatchCode(ctx, batchCode)
	if err != nil || batch == nil {
		return nil, catalog.ErrBatchNotFound
	}

	if batch.IsFull() {
		return nil, catalog.ErrBatchFull
	}

	if !batch.IsPublished() {
		return nil, catalog.ErrBatchNotOpen
	}

	now := time.Now()

	b := &Booking{
		ID:           s.generateID(),
		Code:         s.generateCode(),
		Status:       StatusReserved,
		ReservedAt:   now,
		ExpiredAt:    now.Add(15 * time.Minute),
		CustomerName: customerName,
		Batch:        batch,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.bookingStore.Save(ctx, b); err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Service) generateID() string {
	return uuid.NewString()
}

func (s *Service) generateCode() string {
	now := time.Now()
	date := now.Format("20060102")

	return fmt.Sprintf(
		"BOOK-%s-%s",
		date,
		func() string {
			const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			var sb strings.Builder
			for i := 0; i < 4; i++ {
				sb.WriteByte(letters[rand.Intn(len(letters))])
			}
			return sb.String()
		}(),
	)
}
