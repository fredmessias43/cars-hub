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

type CarModelHandler struct {
	DB *gorm.DB
}

func (h *CarModelHandler) Index(c *gin.Context) {
	car_models := []models.CarModel{}
	config.DB.Preload("Brand").Find(&car_models)

	brands := []models.Brand{}
	config.DB.Find(&brands)

	c.HTML(http.StatusOK, "", templates.CarModelsIndexPage("CarModels page", car_models, brands))
}

func (h *CarModelHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car_model")))

	car_model := models.CarModel{}
	_ = config.DB.Preload("Brand").Find(&car_model, ID)

	c.HTML(http.StatusOK, "", templates.CarModelIndexCard(car_model))
}

func (h *CarModelHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car_model")))

	car_model := models.CarModel{}
	_ = config.DB.Preload("Brand").Find(&car_model, ID)

	brands := []models.Brand{}
	config.DB.Find(&brands)

	c.HTML(http.StatusOK, "", templates.CarModelUpsertForm(car_model, brands))
}

func (h *CarModelHandler) Create(c *gin.Context) {
	car_model := models.CarModel{}

	brands := []models.Brand{}
	config.DB.Find(&brands)

	c.HTML(http.StatusOK, "", templates.CarModelUpsertForm(car_model, brands))
}

func (h *CarModelHandler) Store(c *gin.Context) {
	var car_model models.CarModel

	if err := c.Bind(&car_model); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	_ = config.DB.Create(&car_model)
	_ = config.DB.Preload("Brand").Find(&car_model)

	c.HTML(http.StatusOK, "", templates.CarModelIndexCard(car_model))
}

func (h *CarModelHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car_model")))

	car_model := models.CarModel{ID: ID}

	if err := c.Bind(&car_model); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	config.DB.Model(&car_model).Where("ID = ?", ID).Updates(&car_model)
	_ = config.DB.Preload("Brand").Find(&car_model, ID)

	c.HTML(http.StatusOK, "", templates.CarModelIndexCard(car_model))
}

func (h *CarModelHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("car_model"))

	config.DB.Delete(&models.CarModel{}, ID)

	c.HTML(http.StatusNoContent, "", templates.NoContent())
}
