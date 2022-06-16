package database

import (
	"os"

	"github.com/lucasd-coder/classroom/internal/pkg/logger"
	"github.com/lucasd-coder/classroom/internal/pkg/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := os.Getenv("DATABASE_URL")

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal(err.Error())
	} else {
		logger.Log.Infoln("Postgres Connected")
	}

	db = database

	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
