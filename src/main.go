package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/handlers"
	"github.com/fredmessias43/car-hub/src/models"
	"github.com/gertd/go-pluralize"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := configDb()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	e := echo.New()
	// e.Renderer = &EchoTemplRenderer{}
	e.Static("/assets", "./src/public/assets")

	contactHandler := handlers.ContactHandler{DB: db}
	routerResource(e, "contact", &contactHandler)

	// manufacturerHandler := handlers.ManufacturerHandler{DB: db}
	// routerResource(e, "manufacturer", &manufacturerHandler)

	// brandHandler := handlers.BrandHandler{DB: db}
	// routerResource(e, "brand", &brandHandler)

	// carModelHandler := handlers.CarModelHandler{DB: db}
	// routerResource(e, "car_model", &carModelHandler)

	// carModelVersionHandler := handlers.CarModelVersionHandler{DB: db}
	// routerResource(e, "car_model_version", &carModelVersionHandler)

	// carHandler := handlers.CarHandler{DB: db}
	// routerResource(e, "car", &carHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func routerResource(e *echo.Echo, key string, handler contracts.Handler) {
	singular := pluralize.NewClient().Singular(key)
	plural := pluralize.NewClient().Plural(key)

	e.GET("/"+plural+"", handler.Index)
	e.GET("/"+plural+"/:"+singular+"", handler.Show)
	e.GET("/"+plural+"/:"+singular+"/edit", handler.Edit)
	e.GET("/"+plural+"/create", handler.Create)
	e.POST("/"+plural+"", handler.Store)
	e.PUT("/"+plural+"/:"+singular+"", handler.Update)
	e.DELETE("/"+plural+"/:"+singular+"", handler.Delete)
}

func configDb() (db *gorm.DB, err error) {
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
