package handlers

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/fredmessias43/car-hub/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ContactHandler struct {
	DB *gorm.DB
}

func (h *ContactHandler) Index(c echo.Context) error {
	contacts := []models.Contact{}
	h.DB.Find(&contacts)

	return utils.RenderTemplToHTML(c, templates.ContactsIndexPage("Contacts page", contacts))
}

func (h *ContactHandler) Show(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	return utils.RenderTemplToHTML(c, templates.ContactIndexCard(contact))
}

func (h *ContactHandler) Edit(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	return utils.RenderTemplToHTML(c, templates.ContactUpsertForm(contact))
}

func (h *ContactHandler) Create(c echo.Context) error {
	contact := models.Contact{}
	return utils.RenderTemplToHTML(c, templates.ContactUpsertForm(contact))
}

func (h *ContactHandler) Store(c echo.Context) error {
	var contact models.Contact

	if err := c.Bind(&contact); err != nil {
		return err
	}

	_ = h.DB.Create(&contact)

	return utils.RenderTemplToHTML(c, templates.ContactIndexCard(contact))
}

func (h *ContactHandler) Update(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	if err := c.Bind(&contact); err != nil {
		return err
	}

	h.DB.Model(&contact).Where("ID = ?", ID).Updates(&contact)

	return utils.RenderTemplToHTML(c, templates.ContactIndexCard(contact))
}

func (h *ContactHandler) Delete(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("contact"))

	h.DB.Delete(&models.Contact{}, ID)

	return nil
}
