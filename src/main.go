package main

import (
	"net/http"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/database"
	"github.com/fredmessias43/car-hub/src/handlers"
	"github.com/fredmessias43/car-hub/src/middleware"
	"github.com/fredmessias43/car-hub/src/renderer"
	"github.com/fredmessias43/car-hub/src/templates"
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
	config.Router.Use(config.SetContext())
	config.Router.HTMLRender = &renderer.TemplRender{}
	config.Router.Static("/assets", "./src/public/assets")

	config.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", templates.HomePage())
	})

	/* config.WS = websocket.NewWebsocketServer()
	go config.WS.Run()

	config.Router.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(config.WS, c.Writer, c.Request)
	}) */

	// Pages
	authHandler := handlers.AuthHandler{}
	config.Router.GET("/login", authHandler.ShowLogin)
	config.Router.POST("/login", authHandler.Login)
	config.Router.GET("/register", authHandler.ShowRegister)
	config.Router.POST("/register", authHandler.Register)
	config.Router.POST("/logout", authHandler.Logout)

	// userHandler := handlers.UserHandler{}
	// routerResource(config.Router, "user", &userHandler)

	// auth
	authenticated := config.Router.Group("/")
	authenticated.Use(middleware.AuthMiddleware())
	{
		manufacturerHandler := handlers.ManufacturerHandler{}
		routerResource(authenticated, "manufacturer", &manufacturerHandler)

		brandHandler := handlers.BrandHandler{}
		routerResource(authenticated, "brand", &brandHandler)

		carModelHandler := handlers.CarModelHandler{}
		routerResource(authenticated, "car_model", &carModelHandler)

		carModelVersionHandler := handlers.CarModelVersionHandler{}
		routerResource(authenticated, "car_model_version", &carModelVersionHandler)

		carHandler := handlers.CarHandler{}
		routerResource(authenticated, "car", &carHandler)

		// Other Pages

		integratedHandler := handlers.IntegratedHandler{}
		authenticated.GET("/integrated", integratedHandler.Create)
		authenticated.POST("/integrated", integratedHandler.Store)
		authenticated.GET("/integrated/:car/edit", integratedHandler.Edit)
		authenticated.PUT("/integrated/:car", integratedHandler.Update)

		// Components

		manufacturerComponentHandler := handlers.ManufacturerComponentHandler{}
		authenticated.GET("/manufacturers/:manufacturer/components/option", manufacturerComponentHandler.ShowOptionComponent)
	}

	config.Router.Run()
}

func routerResource(router *gin.RouterGroup, key string, handler contracts.Handler) {
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
