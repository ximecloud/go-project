package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/internal/core/domain"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func New(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Host, config.Username, config.Password, config.Database, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(
		&domain.User{},
	)
	return db, nil
}
