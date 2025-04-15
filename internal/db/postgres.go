package db

import (
	"campaign_engine/internal/models"
	"campaign_engine/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	db, err := gorm.Open(postgres.Open(utils.PostgresDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Campaign{}, &models.Bid{})
	DB = db
}
