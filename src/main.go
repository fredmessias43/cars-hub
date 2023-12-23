package main

import (
	"net/http"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/database"
	"github.com/fredmessias43/car-hub/src/handlers"
	"github.com/fredmessias43/car-hub/src/renderer"
	"github.com/fredmessias43/car-hub/src/websocket"
	"github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	config.DB, err = database.NewDatabaseClient()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	config.Router = gin.Default()
	config.Router.HTMLRender = &renderer.TemplRender{}
	config.Router.Static("/assets", "./src/public/assets")

	config.Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/cars")
	})

	config.WS = websocket.NewWebsocketServer()
	go config.WS.Run()

	config.Router.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(config.WS, c.Writer, c.Request)
	})

	contactHandler := handlers.ContactHandler{}
	routerResource(config.Router, "contact", &contactHandler)

	manufacturerHandler := handlers.ManufacturerHandler{}
	routerResource(config.Router, "manufacturer", &manufacturerHandler)

	brandHandler := handlers.BrandHandler{}
	routerResource(config.Router, "brand", &brandHandler)

	carModelHandler := handlers.CarModelHandler{}
	routerResource(config.Router, "car_model", &carModelHandler)

	carModelVersionHandler := handlers.CarModelVersionHandler{}
	routerResource(config.Router, "car_model_version", &carModelVersionHandler)

	carHandler := handlers.CarHandler{}
	routerResource(config.Router, "car", &carHandler)

	config.Router.Run()
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
