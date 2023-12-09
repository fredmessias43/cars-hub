package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactHandler struct {
	DB *gorm.DB
}

func (h *ContactHandler) Index(c *gin.Context) {
	contacts := []models.Contact{}
	h.DB.Find(&contacts)

	c.HTML(http.StatusOK, "", templates.ContactsIndexPage("Contacts page", contacts))
}

func (h *ContactHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	c.HTML(http.StatusOK, "", templates.ContactIndexCard(contact))
}

func (h *ContactHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	c.HTML(http.StatusOK, "", templates.ContactUpsertForm(contact))
}

func (h *ContactHandler) Create(c *gin.Context) {
	contact := models.Contact{}
	c.HTML(http.StatusOK, "", templates.ContactUpsertForm(contact))
}

func (h *ContactHandler) Store(c *gin.Context) {
	var contact models.Contact

	if err := c.Bind(&contact); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	_ = h.DB.Create(&contact)

	c.HTML(http.StatusOK, "", templates.ContactIndexCard(contact))
}

func (h *ContactHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	if err := c.Bind(&contact); err != nil {
		c.HTML(http.StatusBadRequest, "", err)
		return
	}

	h.DB.Model(&contact).Where("ID = ?", ID).Updates(&contact)

	c.HTML(http.StatusOK, "", templates.ContactIndexCard(contact))
}

func (h *ContactHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("contact"))

	h.DB.Delete(&models.Contact{}, ID)

	c.HTML(http.StatusOK, "", templates.NoContent())
}
