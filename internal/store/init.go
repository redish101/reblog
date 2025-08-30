package store

import (
	"github.com/redish101/reblog/internal/env"
	"github.com/redish101/reblog/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var err error

	if env.Dev {
		DB, err = gorm.Open(sqlite.Open(env.DatabaseURL))
	} else {
		DB, err = gorm.Open(postgres.Open(env.DatabaseURL))
	}

	if err != nil {
		return err
	}

	if err := migrateModels(DB); err != nil {
		return err
	}

	return nil
}

func migrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.CategoryModel{},
		&model.TagModel{},
		&model.PostModel{},
	)
}
