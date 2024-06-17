package config

import (
	"github.com/fredmessias43/car-hub/src/websocket"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

var DB *gorm.DB
var WS *websocket.WsServer
var Router *gin.Engine
var Context *gin.Context
var Store = sessions.NewCookieStore([]byte("secret"))

func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		Context = c
		c.Next()
	}
}

func GetSession() *sessions.Session {
	session, _ := Store.Get(Context.Request, "session")
	return session
}

func SetSession(key string, value interface{}) {
	session := GetSession()
	session.Values[key] = value
	session.Save(Context.Request, Context.Writer)
}
