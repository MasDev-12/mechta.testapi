package db_context

import (
	"fmt"
	"github.com/MasDev-12/mechta.testapi/config"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type DbContext struct {
	Db *gorm.DB
}

func NewDbContext(settings *config.DbSetting) *DbContext {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		settings.Host, settings.User, settings.Password, settings.DBName, settings.Port, settings.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sqlDB from GORM: %v", err)
	}
	sqlDb.SetMaxOpenConns(settings.MaxConnections)
	sqlDb.SetMaxIdleConns(settings.MaxIdleConnections)
	sqlDb.SetConnMaxLifetime(time.Duration(settings.ConnectionMaxLifetime) * time.Second)
	sqlDb.SetConnMaxIdleTime(time.Duration(settings.ConnectionMaxIdleTime) * time.Second)

	if err := db.AutoMigrate(&entities.User{}, &entities.URL{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return &DbContext{Db: db}
}

func (db *DbContext) ClearDatabaseAfterTests() {
	db.Db.Exec("DELETE FROM urls")
	db.Db.Exec("DELETE FROM users")
}
