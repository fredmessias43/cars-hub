package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fredmessias43/car-hub/contracts"
	"github.com/fredmessias43/car-hub/handlers"
	"github.com/fredmessias43/car-hub/models"
	"github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := configDb()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	contactHandler := handlers.ContactHandler{DB: db}
	routerResource(router, "contact", &contactHandler)

	router.Run(":8080")
}

func routerResource(router *gin.Engine, key string, handler contracts.Handler) {
	singular := pluralize.NewClient().Singular(key)
	plural := pluralize.NewClient().Plural(key)

	router.GET("/"+plural+"", handler.Index)
	router.GET("/"+plural+"/:"+singular+"", handler.Show)
	router.GET("/"+plural+"/:"+singular+"/edit", handler.Edit)
	router.GET("/"+plural+"/create", handler.Create)
	router.POST("/"+plural+"", handler.Store)
	router.PUT("/"+plural+"/:"+singular+"", handler.Update)
	router.DELETE("/"+plural+"/:"+singular+"", handler.Delete)
}

func configDb() (db *gorm.DB, err error) {
	dbLog, _ := os.Create("./logs/db.log")
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

	db, err = gorm.Open(sqlite.Open("./database/db.sqlite"), &gorm.Config{
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
