package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BrandHandler struct {
	DB *gorm.DB
}

func (h *BrandHandler) Index(c *gin.Context) {
	brands := []models.Brand{}
	h.DB.Find(&brands)

	c.HTML(http.StatusOK, "pages/brands/index", gin.H{
		"Title":  "Brands page",
		"Brands": brands,
	})
}

func (h *BrandHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("brand"))

	brand := models.Brand{}
	_ = h.DB.Find(&brand, ID)

	c.HTML(http.StatusOK, "partials/brands/index-card", gin.H{
		"Brand": brand.ToMap(),
	})
}

func (h *BrandHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("brand"))

	brand := models.Brand{}
	_ = h.DB.Find(&brand, ID)

	c.HTML(http.StatusOK, "partials/brands/upsert-form", gin.H{
		"Brand": brand.ToMap(),
	})
}

func (h *BrandHandler) Create(c *gin.Context) {
	brand := models.Brand{}
	c.HTML(http.StatusOK, "partials/brands/upsert-form", gin.H{
		"Brand": brand.ToMap(),
	})
}

func (h *BrandHandler) Store(c *gin.Context) {
	var brand models.Brand

	if err := c.ShouldBind(&brand); err != nil {
		return
	}

	_ = h.DB.Find(&brand.Manufacturer, brand.ManufacturerID)

	_ = h.DB.Create(&brand)

	c.HTML(http.StatusOK, "partials/brands/index-card", gin.H{
		"Brand": brand.ToMap(),
	})
}

func (h *BrandHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("brand"))

	brand := models.Brand{}
	_ = h.DB.Find(&brand, ID)

	if err := c.ShouldBind(&brand); err != nil {
		return
	}

	h.DB.Model(&brand).Where("ID = ?", ID).Updates(&brand)

	c.HTML(http.StatusOK, "partials/brands/index-card", gin.H{
		"Brand": brand.ToMap(),
	})
}

func (h *BrandHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("brand"))

	h.DB.Delete(&models.Brand{}, ID)

	return
}
