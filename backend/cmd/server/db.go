package main

import (
	"fmt"

	"lawlens-g/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBSSL,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Contract{},
		&models.Clause{},
		&models.ClauseTrigger{},
		&models.Obligation{},
	)
}
