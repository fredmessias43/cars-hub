package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CarHandler struct {
	DB *gorm.DB
}

func (h *CarHandler) Index(c *gin.Context) {
	cars := []models.Car{}
	h.DB.Find(&cars)

	c.HTML(http.StatusOK, "pages/cars/index", gin.H{
		"Title": "Cars page",
		"Cars":  cars,
	})
}

func (h *CarHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car"))

	car := models.Car{}
	_ = h.DB.Find(&car, ID)

	c.HTML(http.StatusOK, "partials/cars/index-card", gin.H{
		"Car": car.ToMap(),
	})
}

func (h *CarHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car"))

	car := models.Car{}
	_ = h.DB.Find(&car, ID)

	c.HTML(http.StatusOK, "partials/cars/upsert-form", gin.H{
		"Car": car.ToMap(),
	})
}

func (h *CarHandler) Create(c *gin.Context) {
	car := models.Car{}
	c.HTML(http.StatusOK, "partials/cars/upsert-form", gin.H{
		"Car": car.ToMap(),
	})
}

func (h *CarHandler) Store(c *gin.Context) {
	var car models.Car

	if err := c.ShouldBind(&car); err != nil {
		return
	}

	_ = h.DB.Find(&car.CarModelVersion, car.ModelVersionID)

	_ = h.DB.Create(&car)

	c.HTML(http.StatusOK, "partials/cars/index-card", gin.H{
		"Car": car.ToMap(),
	})
}

func (h *CarHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car"))

	car := models.Car{}
	_ = h.DB.Find(&car, ID)

	if err := c.ShouldBind(&car); err != nil {
		return
	}

	h.DB.Model(&car).Where("ID = ?", ID).Updates(&car)

	c.HTML(http.StatusOK, "partials/cars/index-card", gin.H{
		"Car": car.ToMap(),
	})
}

func (h *CarHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car"))

	h.DB.Delete(&models.Car{}, ID)

	return
}
