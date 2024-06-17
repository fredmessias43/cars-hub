package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (h AuthHandler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "", templates.LoginPage())
}

func (h AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("Email")
	password := c.PostForm("Password")

	user := models.GetUser(email, password)
	if user == nil {
		c.HTML(http.StatusUnauthorized, "", templates.Unauthorized(errors.New("invalid email or password")))
		return
	}

	session := config.GetSession()
	session.Values["authenticated"] = true
	session.Values["user"] = strconv.Itoa(user.ID) + "|" + user.Email
	session.Save(c.Request, c.Writer)

	c.Header("HX-Redirect", "/cars")
	c.Redirect(http.StatusSeeOther, "/cars")
}

func (h AuthHandler) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "", templates.RegisterPage())
}

func (h AuthHandler) Register(c *gin.Context) {
	// firstName := c.Param("FirstName")
	// lastName := c.Param("LastName")
	// email := c.Param("Email")
	password := c.Param("Password")
	confirmPassword := c.Param("ConfirmPassword")

	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "", "Passwords do not match")
		return
	}

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", templates.UnprocessableEntity(err))
		return
	}

	config.DB.Create(&user)

	c.Header("HX-Redirect", "/login")
	c.Redirect(http.StatusSeeOther, "/login")
}

func (h AuthHandler) Logout(c *gin.Context) {
	session := config.GetSession()
	session.Values["authenticated"] = false
	session.Values["user"] = ""
	session.Save(c.Request, c.Writer)

	c.Header("HX-Redirect", "/login")
	c.Redirect(http.StatusSeeOther, "/login")
}

func (h AuthHandler) ForgotPassword(c *gin.Context) {
	//
}
