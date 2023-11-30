package handlers

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/fredmessias43/car-hub/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CarModelVersionHandler struct {
	DB *gorm.DB
}

func (h *CarModelVersionHandler) Index(c echo.Context) error {
	carModelVersions := []models.CarModelVersion{}
	h.DB.Find(&carModelVersions)

	return utils.RenderTemplToHTML(c, templates.CarModelVersionsIndexPage("CarModelVersions page", carModelVersions))
}

func (h *CarModelVersionHandler) Show(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("carModelVersion")))

	carModelVersion := models.CarModelVersion{}
	_ = h.DB.Find(&carModelVersion, ID)

	return utils.RenderTemplToHTML(c, templates.CarModelVersionIndexCard(carModelVersion))
}

func (h *CarModelVersionHandler) Edit(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("carModelVersion")))

	carModelVersion := models.CarModelVersion{}
	_ = h.DB.Find(&carModelVersion, ID)

	return utils.RenderTemplToHTML(c, templates.CarModelVersionUpsertForm(carModelVersion))
}

func (h *CarModelVersionHandler) Create(c echo.Context) error {
	carModelVersion := models.CarModelVersion{}
	return utils.RenderTemplToHTML(c, templates.CarModelVersionUpsertForm(carModelVersion))
}

func (h *CarModelVersionHandler) Store(c echo.Context) error {
	var carModelVersion models.CarModelVersion

	if err := c.Bind(&carModelVersion); err != nil {
		return err
	}

	_ = h.DB.Create(&carModelVersion)

	return utils.RenderTemplToHTML(c, templates.CarModelVersionIndexCard(carModelVersion))
}

func (h *CarModelVersionHandler) Update(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("carModelVersion")))

	carModelVersion := models.CarModelVersion{}
	_ = h.DB.Find(&carModelVersion, ID)

	if err := c.Bind(&carModelVersion); err != nil {
		return err
	}

	h.DB.Model(&carModelVersion).Where("ID = ?", ID).Updates(&carModelVersion)

	return utils.RenderTemplToHTML(c, templates.CarModelVersionIndexCard(carModelVersion))
}

func (h *CarModelVersionHandler) Delete(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("carModelVersion"))

	h.DB.Delete(&models.CarModelVersion{}, ID)

	return nil
}
