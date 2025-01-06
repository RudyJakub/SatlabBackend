package server

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

type templateRegistry struct {
	templates map[string]*template.Template
}

func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)

		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func Renderer() *templateRegistry {
	templates := make(map[string]*template.Template)

	templates["home.html"] = template.Must(template.ParseFiles("web/templates/home.html", "web/templates/base.html"))
	templates["edit_article.html"] = template.Must(template.ParseFiles("web/templates/edit_article.html", "web/templates/base.html"))
	templates["login.html"] = template.Must(template.ParseFiles("web/templates/login.html"))
	templates["logout.html"] = template.Must(template.ParseFiles("web/templates/logout.html"))

	return &templateRegistry{
		templates: templates,
	}
}
