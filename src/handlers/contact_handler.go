package handlers

import (
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactHandler struct {
	DB *gorm.DB
}

func (h *ContactHandler) Index(c *gin.Context) {
	contacts := []models.Contact{}
	h.DB.Find(&contacts)

	c.HTML(http.StatusOK, "pages/contacts/index", gin.H{
		"title":    "Contacts page",
		"contacts": contacts,
	})
}

func (h *ContactHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	c.HTML(http.StatusOK, "partials/contacts/index-card", contact.ToMap())
}

func (h *ContactHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	c.HTML(http.StatusOK, "partials/contacts/upsert-form", contact.ToMap())
}

func (h *ContactHandler) Create(c *gin.Context) {
	contact := models.Contact{}
	c.HTML(http.StatusOK, "partials/contacts/upsert-form", contact.ToMap())
}

func (h *ContactHandler) Store(c *gin.Context) {
	var contact models.Contact

	if err := c.ShouldBind(&contact); err != nil {
		return
	}

	_ = h.DB.Create(&contact)

	c.HTML(http.StatusOK, "partials/contacts/index-card", contact.ToMap())
}

func (h *ContactHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("contact")))

	contact := models.Contact{}
	_ = h.DB.Find(&contact, ID)

	if err := c.ShouldBind(&contact); err != nil {
		return
	}

	h.DB.Model(&contact).Where("ID = ?", ID).Updates(&contact)

	c.HTML(http.StatusOK, "partials/contacts/index-card", contact.ToMap())
}

func (h *ContactHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("contact"))

	h.DB.Delete(&models.Contact{}, ID)

	return
}
