package booking

import (
	"app/internal/database/entity"
	"context"

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
	e := ToBookingEntity(booking)

	err := s.DB.WithContext(ctx).
		Model(&entity.Booking{}).
		Create(e).
		Error
	if err != nil {
		return err
	}

	return nil
}
