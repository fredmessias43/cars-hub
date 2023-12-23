package database

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fredmessias43/car-hub/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseClient() (db *gorm.DB, err error) {
	dbLog, _ := os.Create("./src/logs/db.log")
	out := io.MultiWriter(dbLog)
	newLogger := logger.New(
		log.New(out, "\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err = gorm.Open(sqlite.Open("./src/database/db.sqlite"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&models.Contact{},
		&models.Manufacturer{},
		&models.Brand{},
		&models.CarModel{},
		&models.CarModelVersion{},
		&models.Car{},
	)

	return db, nil
}
