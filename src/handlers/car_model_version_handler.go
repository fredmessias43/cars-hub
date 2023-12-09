package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarModelVersionHandler struct {
	DB *gorm.DB
}

func (h *CarModelVersionHandler) Index(c *gin.Context) {
	carModelVersions := []models.CarModelVersion{}
	h.DB.Preload("CarModel").Find(&carModelVersions)

	carModels := []models.CarModel{}
	_ = h.DB.Find(&carModels)

	c.HTML(http.StatusOK, "", templates.CarModelVersionsIndexPage("CarModelVersions page", carModelVersions, carModels))
}

func (h *CarModelVersionHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("carModelVersion")))

	carModelVersion := models.CarModelVersion{ID: ID}
	_ = h.DB.Preload("CarModel").Find(&carModelVersion)

	c.HTML(http.StatusOK, "", templates.CarModelVersionIndexCard(carModelVersion))
}

func (h *CarModelVersionHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("carModelVersion")))

	carModelVersion := models.CarModelVersion{ID: ID}
	_ = h.DB.Preload("CarModel").Find(&carModelVersion)

	carModels := []models.CarModel{}
	_ = h.DB.Find(&carModels)

	c.HTML(http.StatusOK, "", templates.CarModelVersionUpsertForm(carModelVersion, carModels))
}

func (h *CarModelVersionHandler) Create(c *gin.Context) {
	carModelVersion := models.CarModelVersion{}

	carModels := []models.CarModel{}
	_ = h.DB.Find(&carModels)

	c.HTML(http.StatusOK, "", templates.CarModelVersionUpsertForm(carModelVersion, carModels))
}

func (h *CarModelVersionHandler) Store(c *gin.Context) {
	var carModelVersion models.CarModelVersion

	if err := c.Bind(&carModelVersion); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	_ = h.DB.Create(&carModelVersion)
	_ = h.DB.Preload("CarModel").Find(&carModelVersion)

	c.HTML(http.StatusOK, "", templates.CarModelVersionIndexCard(carModelVersion))
}

func (h *CarModelVersionHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("carModelVersion")))

	carModelVersion := models.CarModelVersion{ID: ID}

	if err := c.Bind(&carModelVersion); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	h.DB.Model(&carModelVersion).Where("ID = ?", ID).Updates(&carModelVersion)
	_ = h.DB.Preload("CarModel").Find(&carModelVersion)

	c.HTML(http.StatusOK, "", templates.CarModelVersionIndexCard(carModelVersion))
}

func (h *CarModelVersionHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("carModelVersion"))

	h.DB.Delete(&models.CarModelVersion{ID: ID})

	c.HTML(http.StatusNoContent, "", templates.NoContent())
}
