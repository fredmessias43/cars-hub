package handlers

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/fredmessias43/car-hub/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BrandHandler struct {
	DB *gorm.DB
}

func (h *BrandHandler) Index(c echo.Context) error {
	brands := []models.Brand{}
	h.DB.Find(&brands)

	return utils.RenderTemplToHTML(c, templates.BrandsIndexPage("Brands page", brands))
}

func (h *BrandHandler) Show(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("brand")))

	brand := models.Brand{}
	_ = h.DB.Find(&brand, ID)

	return utils.RenderTemplToHTML(c, templates.BrandIndexCard(brand))
}

func (h *BrandHandler) Edit(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("brand")))

	brand := models.Brand{}
	_ = h.DB.Find(&brand, ID)

	return utils.RenderTemplToHTML(c, templates.BrandUpsertForm(brand))
}

func (h *BrandHandler) Create(c echo.Context) error {
	brand := models.Brand{}
	return utils.RenderTemplToHTML(c, templates.BrandUpsertForm(brand))
}

func (h *BrandHandler) Store(c echo.Context) error {
	var brand models.Brand

	if err := c.Bind(&brand); err != nil {
		return err
	}

	_ = h.DB.Create(&brand)

	return utils.RenderTemplToHTML(c, templates.BrandIndexCard(brand))
}

func (h *BrandHandler) Update(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("brand")))

	brand := models.Brand{}
	_ = h.DB.Find(&brand, ID)

	if err := c.Bind(&brand); err != nil {
		return err
	}

	h.DB.Model(&brand).Where("ID = ?", ID).Updates(&brand)

	return utils.RenderTemplToHTML(c, templates.BrandIndexCard(brand))
}

func (h *BrandHandler) Delete(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("brand"))

	h.DB.Delete(&models.Brand{}, ID)

	return nil
}
