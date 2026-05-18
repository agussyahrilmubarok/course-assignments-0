package booking

import (
	"app/internal/database/entity"
	"app/internal/logger"
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) Save(ctx context.Context, booking *Booking) error {
	log := logger.FromCtx(ctx)

	log.Info("store save booking",
		zap.String("booking_id", booking.ID),
		zap.String("booking_code", booking.Code),
	)

	e := ToBookingEntity(booking)

	err := s.DB.WithContext(ctx).
		Model(&entity.Booking{}).
		Create(e).
		Error
	if err != nil {
		log.Error("store save booking failed",
			zap.String("booking_id", booking.ID),
			zap.String("booking_code", booking.Code),
			zap.Error(err),
		)
		return err
	}

	log.Info("store save booking success",
		zap.String("booking_id", booking.ID),
		zap.String("booking_code", booking.Code),
	)

	return nil
}
