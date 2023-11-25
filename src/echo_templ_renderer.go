package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type EchoTemplRenderer struct {
	templates *template.Template
}

func (t GinTemplRender) writeContentType(w http.ResponseWriter) {
	w.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
}

func (t *EchoTemplRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
