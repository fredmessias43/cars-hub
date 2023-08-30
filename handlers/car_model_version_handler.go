package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarModelVersionHandler struct {
	DB *gorm.DB
}

func (h *CarModelVersionHandler) Index(c *gin.Context) {
	carModelVersions := []models.CarModelVersion{}
	h.DB.Find(&carModelVersions)

	c.HTML(http.StatusOK, "pages/car-model-versions/index", gin.H{
		"Title":            "CarModelVersions page",
		"CarModelVersions": carModelVersions,
	})
}

func (h *CarModelVersionHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model_version"))

	carModelVersion := models.CarModelVersion{}
	_ = h.DB.Find(&carModelVersion, ID)

	c.HTML(http.StatusOK, "partials/car-model-versions/index-card", gin.H{
		"CarModelVersion": carModelVersion.ToMap(),
	})
}

func (h *CarModelVersionHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model_version"))

	carModelVersion := models.CarModelVersion{}
	_ = h.DB.Find(&carModelVersion, ID)

	c.HTML(http.StatusOK, "partials/car-model-versions/upsert-form", gin.H{
		"CarModelVersion": carModelVersion.ToMap(),
	})
}

func (h *CarModelVersionHandler) Create(c *gin.Context) {
	carModelVersion := models.CarModelVersion{}
	c.HTML(http.StatusOK, "partials/car-model-versions/upsert-form", gin.H{
		"CarModelVersion": carModelVersion.ToMap(),
	})
}

func (h *CarModelVersionHandler) Store(c *gin.Context) {
	var carModelVersion models.CarModelVersion

	if err := c.ShouldBind(&carModelVersion); err != nil {
		return
	}

	_ = h.DB.Find(&carModelVersion.CarModel, carModelVersion.ModelID)

	_ = h.DB.Create(&carModelVersion)

	c.HTML(http.StatusOK, "partials/car-model-versions/index-card", gin.H{
		"CarModelVersion": carModelVersion.ToMap(),
	})
}

func (h *CarModelVersionHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model_version"))

	carModelVersion := models.CarModelVersion{}
	_ = h.DB.Find(&carModelVersion, ID)

	if err := c.ShouldBind(&carModelVersion); err != nil {
		return
	}

	h.DB.Model(&carModelVersion).Where("ID = ?", ID).Updates(&carModelVersion)

	c.HTML(http.StatusOK, "partials/car-model-versions/index-card", gin.H{
		"CarModelVersion": carModelVersion.ToMap(),
	})
}

func (h *CarModelVersionHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model_version"))

	h.DB.Delete(&models.CarModelVersion{}, ID)

	return
}
