package handlers

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/fredmessias43/car-hub/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CarModelHandler struct {
	DB *gorm.DB
}

func (h *CarModelHandler) Index(c echo.Context) error {
	car_models := []models.CarModel{}
	h.DB.Find(&car_models)

	return utils.RenderTemplToHTML(c, templates.CarModelsIndexPage("CarModels page", car_models))
}

func (h *CarModelHandler) Show(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("car_model")))

	car_model := models.CarModel{}
	_ = h.DB.Find(&car_model, ID)

	return utils.RenderTemplToHTML(c, templates.CarModelIndexCard(car_model))
}

func (h *CarModelHandler) Edit(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("car_model")))

	car_model := models.CarModel{}
	_ = h.DB.Find(&car_model, ID)

	return utils.RenderTemplToHTML(c, templates.CarModelUpsertForm(car_model))
}

func (h *CarModelHandler) Create(c echo.Context) error {
	car_model := models.CarModel{}
	return utils.RenderTemplToHTML(c, templates.CarModelUpsertForm(car_model))
}

func (h *CarModelHandler) Store(c echo.Context) error {
	var car_model models.CarModel

	if err := c.Bind(&car_model); err != nil {
		return err
	}

	_ = h.DB.Create(&car_model)

	return utils.RenderTemplToHTML(c, templates.CarModelIndexCard(car_model))
}

func (h *CarModelHandler) Update(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("car_model")))

	car_model := models.CarModel{}
	_ = h.DB.Find(&car_model, ID)

	if err := c.Bind(&car_model); err != nil {
		return err
	}

	h.DB.Model(&car_model).Where("ID = ?", ID).Updates(&car_model)

	return utils.RenderTemplToHTML(c, templates.CarModelIndexCard(car_model))
}

func (h *CarModelHandler) Delete(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("car_model"))

	h.DB.Delete(&models.CarModel{}, ID)

	return nil
}
