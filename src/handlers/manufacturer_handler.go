package handlers

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/fredmessias43/car-hub/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ManufacturerHandler struct {
	DB *gorm.DB
}

func (h *ManufacturerHandler) Index(c echo.Context) error {
	manufacturers := []models.Manufacturer{}
	h.DB.Find(&manufacturers)

	return utils.RenderTemplToHTML(c, templates.ManufacturersIndexPage("Manufacturers page", manufacturers))
}

func (h *ManufacturerHandler) Show(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = h.DB.Find(&manufacturer, ID)

	return utils.RenderTemplToHTML(c, templates.ManufacturerIndexCard(manufacturer))
}

func (h *ManufacturerHandler) Edit(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = h.DB.Find(&manufacturer, ID)

	return utils.RenderTemplToHTML(c, templates.ManufacturerUpsertForm(manufacturer))
}

func (h *ManufacturerHandler) Create(c echo.Context) error {
	manufacturer := models.Manufacturer{}
	return utils.RenderTemplToHTML(c, templates.ManufacturerUpsertForm(manufacturer))
}

func (h *ManufacturerHandler) Store(c echo.Context) error {
	var manufacturer models.Manufacturer

	if err := c.Bind(&manufacturer); err != nil {
		return err
	}

	_ = h.DB.Create(&manufacturer)

	return utils.RenderTemplToHTML(c, templates.ManufacturerIndexCard(manufacturer))
}

func (h *ManufacturerHandler) Update(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("manufacturer")))

	manufacturer := models.Manufacturer{}
	_ = h.DB.Find(&manufacturer, ID)

	if err := c.Bind(&manufacturer); err != nil {
		return err
	}

	h.DB.Model(&manufacturer).Where("ID = ?", ID).Updates(&manufacturer)

	return utils.RenderTemplToHTML(c, templates.ManufacturerIndexCard(manufacturer))
}

func (h *ManufacturerHandler) Delete(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("manufacturer"))

	h.DB.Delete(&models.Manufacturer{}, ID)

	return nil
}
