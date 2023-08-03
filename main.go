package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fredmessias43/car-hub/handlers"
	"github.com/fredmessias43/car-hub/models"
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

	//

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	contactHandler := handlers.ContactHandler{DB: db}
	router.GET("/", contactHandler.Index)
	router.GET("/contact/:contact", contactHandler.Show)
	router.GET("/contact/:contact/edit", contactHandler.Edit)
	// router.GET("/contact/create", contactHandler.Create)
	router.POST("/contact", contactHandler.Store)
	router.PUT("/contact/:contact", contactHandler.Update)
	router.DELETE("/contact/:contact", contactHandler.Delete)

	router.Run(":8080")
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
	db.AutoMigrate(&models.Contact{})

	return db, nil
}
