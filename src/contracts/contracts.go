package contracts

import "github.com/gin-gonic/gin"

type Handler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Edit(c *gin.Context)
	Create(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Model interface {
	ToMap() map[string]any
	ToJson() []byte
}
