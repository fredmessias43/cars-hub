package contracts

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Index(c echo.Context) error
	Show(c echo.Context) error
	Edit(c echo.Context) error
	Create(c echo.Context) error
	Store(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
