package handlers

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/fredmessias43/car-hub/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CarHandler struct {
	DB *gorm.DB
}

func (h *CarHandler) Index(c echo.Context) error {
	cars := []models.Car{}
	h.DB.Find(&cars)

	return utils.RenderTemplToHTML(c, templates.CarsIndexPage("Cars page", cars))
}

func (h *CarHandler) Show(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("car")))

	car := models.Car{}
	_ = h.DB.Find(&car, ID)

	return utils.RenderTemplToHTML(c, templates.CarIndexCard(car))
}

func (h *CarHandler) Edit(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("car")))

	car := models.Car{}
	_ = h.DB.Find(&car, ID)

	return utils.RenderTemplToHTML(c, templates.CarUpsertForm(car))
}

func (h *CarHandler) Create(c echo.Context) error {
	car := models.Car{}
	return utils.RenderTemplToHTML(c, templates.CarUpsertForm(car))
}

func (h *CarHandler) Store(c echo.Context) error {
	var car models.Car

	if err := c.Bind(&car); err != nil {
		return err
	}

	_ = h.DB.Create(&car)

	return utils.RenderTemplToHTML(c, templates.CarIndexCard(car))
}

func (h *CarHandler) Update(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("car")))

	car := models.Car{}
	_ = h.DB.Find(&car, ID)

	if err := c.Bind(&car); err != nil {
		return err
	}

	h.DB.Model(&car).Where("ID = ?", ID).Updates(&car)

	return utils.RenderTemplToHTML(c, templates.CarIndexCard(car))
}

func (h *CarHandler) Delete(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("car"))

	h.DB.Delete(&models.Car{}, ID)

	return nil
}
