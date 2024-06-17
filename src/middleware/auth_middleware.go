package middleware

import (
	"net/http"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := config.GetSession()
		if session.Values["authenticated"] != true {
			c.Header("HX-Redirect", "/login")
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}
		c.Next()
	}
}
