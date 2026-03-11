package catalog

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

func (s *Store) FindBatchCourseByBatchCode(ctx context.Context, batchCode string) (*Batch, error) {
	var e entity.Batch

	err := s.DB.WithContext(ctx).
		Model(&entity.Batch{}).
		Preload("Course").
		Where("code = ?", batchCode).
		First(&e).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrBatchNotFound
		}
		return nil, err
	}

	return ToBatch(&e), nil
}
