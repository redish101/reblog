package store

import (
	"git.liteyuki.org/redish101/reblog/internal/env"
	"git.liteyuki.org/redish101/reblog/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error

	if env.Dev {
		db, err = gorm.Open(sqlite.Open(env.DatabaseURL))
	} else {
		db, err = gorm.Open(postgres.Open(env.DatabaseURL))
	}

	if err != nil {
		return err
	}

	if err := migrateModels(db); err != nil {
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
