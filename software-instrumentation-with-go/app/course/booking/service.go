package booking

import (
	"app/course/catalog"
	"app/internal/logger"
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
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
	log := logger.FromCtx(ctx)

	log.Info("create booking start",
		zap.String("batch_code", batchCode),
		zap.String("customer_name", customerName),
	)

	if customerName == "" {
		log.Warn("invalid customer name")
		return nil, ErrInvalidCustomerName
	}

	batch, err := s.catalogStore.FindBatchCourseByBatchCode(ctx, batchCode)
	if err != nil || batch == nil {
		log.Warn("batch not found",
			zap.String("batch_code", batchCode),
			zap.Error(err),
		)
		return nil, catalog.ErrBatchNotFound
	}

	if batch.IsFull() {
		log.Warn("batch full",
			zap.String("batch_code", batchCode),
		)
		return nil, catalog.ErrBatchFull
	}

	if !batch.IsPublished() {
		log.Warn("batch not open",
			zap.String("batch_code", batchCode),
		)
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
		log.Error("failed save booking",
			zap.String("batch_code", batchCode),
			zap.String("booking_code", b.Code),
			zap.Error(err),
		)
		return nil, err
	}

	log.Info("booking created",
		zap.String("booking_id", b.ID),
		zap.String("booking_code", b.Code),
		zap.String("batch_code", batchCode),
		zap.String("customer_name", customerName),
	)

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
