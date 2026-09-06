// db/database.go
package db

import (
	"fmt"
	"log"

	"github.com/nwangui-dev/SchoolMngSystem/configs"
	"github.com/nwangui-dev/SchoolMngSystem/store"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormDB() {
	cfg := configs.AppConfig.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	var err error
	store.GDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Enable UUID extension
	if err := store.GDB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err)
	}

	// Auto-migrate models
	err = store.GDB.AutoMigrate(Tables...)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
}