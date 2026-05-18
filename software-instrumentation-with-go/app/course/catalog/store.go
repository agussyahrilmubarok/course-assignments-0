package catalog

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

func (s *Store) FindBatchCourseByBatchCode(ctx context.Context, batchCode string) (*Batch, error) {
	log := logger.FromCtx(ctx)

	log.Info("find batch by code",
		zap.String("batch_code", batchCode),
	)

	var e entity.Batch

	err := s.DB.WithContext(ctx).
		Model(&entity.Batch{}).
		Preload("Course").
		Where("code = ?", batchCode).
		First(&e).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warn("batch not found",
				zap.String("batch_code", batchCode),
			)
			return nil, ErrBatchNotFound
		}
		log.Error("db error find batch",
			zap.String("batch_code", batchCode),
			zap.Error(err),
		)
		return nil, err
	}

	log.Info("batch found",
		zap.String("batch_code", batchCode),
		zap.String("course_id", e.CourseID),
	)

	return ToBatch(&e), nil
}
