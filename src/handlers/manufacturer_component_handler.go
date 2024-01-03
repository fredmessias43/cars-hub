package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ManufacturerComponentHandler struct {
	DB *gorm.DB
}

func (h ManufacturerComponentHandler) ShowOptionComponent(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = config.DB.Find(&manufacturer, ID)

	c.HTML(http.StatusOK, "", templates.SelectOption(manufacturer.Name, strconv.Itoa(manufacturer.ID), false))
}
