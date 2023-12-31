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

type CarHandler struct {
	DB *gorm.DB
}

func (h CarHandler) Index(c *gin.Context) {
	cars := []models.Car{}
	config.DB.Preload("CarModelVersion").Find(&cars)

	carModelVersions := []models.CarModelVersion{}
	_ = config.DB.Find(&carModelVersions)

	c.HTML(http.StatusOK, "", templates.CarsIndexPage("Cars page", cars, carModelVersions))
}

func (h CarHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car")))

	car := models.Car{ID: ID}
	_ = config.DB.Preload("CarModelVersion").Find(&car)

	c.HTML(http.StatusOK, "", templates.CarIndexCard(car))
}

func (h CarHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car")))

	car := models.Car{ID: ID}
	_ = config.DB.Preload("CarModelVersion").Find(&car)

	carModelVersions := []models.CarModelVersion{}
	_ = config.DB.Find(&carModelVersions)

	c.HTML(http.StatusOK, "", templates.CarUpsertForm(car, carModelVersions))
}

func (h CarHandler) Create(c *gin.Context) {
	car := models.Car{}

	carModelVersions := []models.CarModelVersion{}
	_ = config.DB.Find(&carModelVersions)

	c.HTML(http.StatusOK, "", templates.CarUpsertForm(car, carModelVersions))
}

func (h CarHandler) Store(c *gin.Context) {
	var car models.Car

	if err := c.Bind(&car); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	_ = config.DB.Create(&car)
	_ = config.DB.Preload("CarModelVersion").Find(&car)

	c.HTML(http.StatusOK, "", templates.CarIndexCard(car))
}

func (h CarHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car")))

	car := models.Car{ID: ID}

	if err := c.Bind(&car); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	config.DB.Model(&car).Where("ID = ?", ID).Updates(&car)
	_ = config.DB.Preload("CarModelVersion").Find(&car)

	c.HTML(http.StatusOK, "", templates.CarIndexCard(car))
}

func (h CarHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car"))

	config.DB.Delete(&models.Car{ID: ID})

	c.HTML(http.StatusNoContent, "", templates.NoContent())
}
