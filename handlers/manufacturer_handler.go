package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ManufacturerHandler struct {
	DB *gorm.DB
}

func (h *ManufacturerHandler) Index(c *gin.Context) {
	manufacturers := []models.Manufacturer{}
	h.DB.Find(&manufacturers)

	c.HTML(http.StatusOK, "pages/manufacturers/index", gin.H{
		"Title":         "Manufacturers page",
		"Manufacturers": manufacturers,
	})
}

func (h *ManufacturerHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("manufacturer"))

	manufacturer := models.Manufacturer{}
	_ = h.DB.Find(&manufacturer, ID)

	c.HTML(http.StatusOK, "partials/manufacturers/index-card", gin.H{
		"Manufacturer": manufacturer.ToMap(),
	})
}

func (h *ManufacturerHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("manufacturer"))

	manufacturer := models.Manufacturer{}
	_ = h.DB.Find(&manufacturer, ID)

	c.HTML(http.StatusOK, "partials/manufacturers/upsert-form", gin.H{
		"Manufacturer": manufacturer.ToMap(),
	})
}

func (h *ManufacturerHandler) Create(c *gin.Context) {
	manufacturer := models.Manufacturer{}
	c.HTML(http.StatusOK, "partials/manufacturers/upsert-form", gin.H{
		"Manufacturer": manufacturer.ToMap(),
	})
}

func (h *ManufacturerHandler) Store(c *gin.Context) {
	var manufacturer models.Manufacturer

	if err := c.ShouldBind(&manufacturer); err != nil {
		return
	}

	_ = h.DB.Create(&manufacturer)

	c.HTML(http.StatusOK, "partials/manufacturers/index-card", gin.H{
		"Manufacturer": manufacturer.ToMap(),
	})
}

func (h *ManufacturerHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("manufacturer"))

	manufacturer := models.Manufacturer{}
	_ = h.DB.Find(&manufacturer, ID)

	if err := c.ShouldBind(&manufacturer); err != nil {
		return
	}

	h.DB.Model(&manufacturer).Where("ID = ?", ID).Updates(&manufacturer)

	c.HTML(http.StatusOK, "partials/manufacturers/index-card", gin.H{
		"Manufacturer": manufacturer.ToMap(),
	})
}

func (h *ManufacturerHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("manufacturer"))

	h.DB.Delete(&models.Manufacturer{}, ID)

	return
}
