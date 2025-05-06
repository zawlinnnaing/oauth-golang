package database

import (
	"fmt"
	"github.com/zawlinnnaing/oauth-golang/authorization-server/modules/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
		config.DB_SSL,
		config.DB_TIME_ZONE,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return DB, err
}
