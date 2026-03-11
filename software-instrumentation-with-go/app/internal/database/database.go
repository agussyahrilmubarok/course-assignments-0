package database

import (
	"app/internal/config"
	"app/internal/database/entity"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	pg := cfg.Postgres

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		pg.Host,
		pg.User,
		pg.Password,
		pg.Name,
		pg.Port,
		pg.SSLMode,
		pg.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(pg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(pg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(pg.ConnMaxLifetime) * time.Minute)

	if err := Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Course{},
		&entity.Batch{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate schema: %w", err)
	}

	return nil
}
