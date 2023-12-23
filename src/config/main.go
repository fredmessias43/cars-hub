package config

import (
	"github.com/fredmessias43/car-hub/src/websocket"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB
var WS *websocket.WsServer
var Router *gin.Engine
