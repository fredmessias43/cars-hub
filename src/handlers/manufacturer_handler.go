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

type ManufacturerHandler struct {
	DB *gorm.DB
}

func (h *ManufacturerHandler) Index(c *gin.Context) {
	manufacturers := []models.Manufacturer{}
	config.DB.Find(&manufacturers)

	c.HTML(http.StatusOK, "", templates.ManufacturersIndexPage("Manufacturers page", manufacturers))
}

func (h *ManufacturerHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = config.DB.Find(&manufacturer, ID)

	c.HTML(http.StatusOK, "", templates.ManufacturerIndexCard(manufacturer))
}

func (h *ManufacturerHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = config.DB.Find(&manufacturer, ID)

	c.HTML(http.StatusOK, "", templates.ManufacturerUpsertForm(manufacturer))
}

func (h *ManufacturerHandler) Create(c *gin.Context) {
	manufacturer := models.Manufacturer{}
	c.HTML(http.StatusOK, "", templates.ManufacturerUpsertForm(manufacturer))
}

func (h *ManufacturerHandler) Store(c *gin.Context) {
	var manufacturer models.Manufacturer

	if err := c.Bind(&manufacturer); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	_ = config.DB.Create(&manufacturer)

	c.HTML(http.StatusOK, "", templates.ManufacturerIndexCard(manufacturer))
}

func (h *ManufacturerHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = config.DB.Find(&manufacturer, ID)

	if err := c.Bind(&manufacturer); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	config.DB.Model(&manufacturer).Where("ID = ?", ID).Updates(&manufacturer)

	c.HTML(http.StatusOK, "", templates.ManufacturerIndexCard(manufacturer))
}

func (h *ManufacturerHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("manufacturer"))

	config.DB.Delete(&models.Manufacturer{}, ID)

	c.HTML(http.StatusOK, "", templates.NoContent())
}
