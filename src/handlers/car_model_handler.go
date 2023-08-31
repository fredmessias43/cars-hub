package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarModelHandler struct {
	DB *gorm.DB
}

func (h *CarModelHandler) Index(c *gin.Context) {
	carModels := []models.CarModel{}
	h.DB.Find(&carModels)

	c.HTML(http.StatusOK, "pages/car-models/index", gin.H{
		"Title":     "CarModels page",
		"CarModels": carModels,
	})
}

func (h *CarModelHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model"))

	carModel := models.CarModel{}
	_ = h.DB.Find(&carModel, ID)

	c.HTML(http.StatusOK, "partials/car-models/index-card", gin.H{
		"CarModel": carModel.ToMap(),
	})
}

func (h *CarModelHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model"))

	carModel := models.CarModel{}
	_ = h.DB.Find(&carModel, ID)

	c.HTML(http.StatusOK, "partials/car-models/upsert-form", gin.H{
		"CarModel": carModel.ToMap(),
	})
}

func (h *CarModelHandler) Create(c *gin.Context) {
	carModel := models.CarModel{}
	c.HTML(http.StatusOK, "partials/car-models/upsert-form", gin.H{
		"CarModel": carModel.ToMap(),
	})
}

func (h *CarModelHandler) Store(c *gin.Context) {
	var carModel models.CarModel

	if err := c.ShouldBind(&carModel); err != nil {
		return
	}

	_ = h.DB.Find(&carModel.Brand, carModel.BrandID)

	_ = h.DB.Create(&carModel)

	c.HTML(http.StatusOK, "partials/car-models/index-card", gin.H{
		"CarModel": carModel.ToMap(),
	})
}

func (h *CarModelHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model"))

	carModel := models.CarModel{}
	_ = h.DB.Find(&carModel, ID)

	if err := c.ShouldBind(&carModel); err != nil {
		return
	}

	h.DB.Model(&carModel).Where("ID = ?", ID).Updates(&carModel)

	c.HTML(http.StatusOK, "partials/car-models/index-card", gin.H{
		"CarModel": carModel.ToMap(),
	})
}

func (h *CarModelHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model"))

	h.DB.Delete(&models.CarModel{}, ID)

	return
}
