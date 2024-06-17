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

type UserHandler struct {
	DB *gorm.DB
}

func (h UserHandler) Index(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)

	c.HTML(http.StatusOK, "", templates.UsersIndexPage("Users page", users))
}

func (h UserHandler) Show(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("user")))

	user := models.User{}
	_ = config.DB.Find(&user, ID)

	c.HTML(http.StatusOK, "", templates.UserIndexCard(user))
}

func (h UserHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("user")))

	user := models.User{}
	_ = config.DB.Find(&user, ID)

	c.HTML(http.StatusOK, "", templates.UserUpsertForm(user))
}

func (h UserHandler) Create(c *gin.Context) {
	user := models.User{}
	c.HTML(http.StatusOK, "", templates.UserUpsertForm(user))
}

func (h UserHandler) Store(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	_ = config.DB.Create(&user)

	c.HTML(http.StatusOK, "", templates.UserIndexCard(user))
}

func (h UserHandler) Update(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("user")))

	user := models.User{}
	_ = config.DB.Find(&user, ID)

	if err := c.Bind(&user); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	config.DB.Model(&user).Where("ID = ?", ID).Updates(&user)

	c.HTML(http.StatusOK, "", templates.UserIndexCard(user))
}

func (h UserHandler) Delete(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("user"))

	config.DB.Delete(&models.User{}, ID)

	c.HTML(http.StatusOK, "", templates.NoContent())
}
