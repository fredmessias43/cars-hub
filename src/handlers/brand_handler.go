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

type BrandHandler struct {
	DB *gorm.DB
}

func (h BrandHandler) Index(c *gin.Context) {
	brands := []models.Brand{}
	config.DB.Preload("Manufacturer").Find(&brands)

	manufacturers := []models.Manufacturer{}
	config.DB.Find(&manufacturers)

	c.HTML(http.StatusOK, "", templates.BrandsIndexPage("Brands page", brands, manufacturers))
}

func (h BrandHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("brand")))

	brand := models.Brand{ID: ID}
	_ = config.DB.Preload("Manufacturer").Find(&brand)

	c.HTML(http.StatusOK, "", templates.BrandIndexCard(brand))
}

func (h BrandHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("brand")))

	brand := models.Brand{}
	_ = config.DB.Preload("Manufacturer").Find(&brand, ID)

	manufacturers := []models.Manufacturer{}
	config.DB.Find(&manufacturers)

	c.HTML(http.StatusOK, "", templates.BrandUpsertForm(brand, manufacturers))
}

func (h BrandHandler) Create(c *gin.Context) {
	brand := models.Brand{}

	manufacturers := []models.Manufacturer{}
	config.DB.Find(&manufacturers)

	c.HTML(http.StatusOK, "", templates.BrandUpsertForm(brand, manufacturers))
}

func (h BrandHandler) Store(c *gin.Context) {
	brand := models.Brand{}

	if err := c.Bind(&brand); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	_ = config.DB.Create(&brand)
	_ = config.DB.Preload("Manufacturer").Find(&brand)

	c.HTML(http.StatusOK, "", templates.BrandIndexCard(brand))
}

func (h BrandHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("brand")))

	brand := models.Brand{ID: ID}
	if err := c.ShouldBind(&brand); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	config.DB.Model(&brand).Where("ID = ?", ID).Updates(&brand)
	_ = config.DB.Preload("Manufacturer").Find(&brand, ID)

	c.HTML(http.StatusOK, "", templates.BrandIndexCard(brand))
}

func (h BrandHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("brand"))

	config.DB.Delete(&models.Brand{ID: ID}, ID)

	c.HTML(http.StatusOK, "", templates.NoContent())
}
