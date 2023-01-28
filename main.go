package main

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	box := packr.New("templates", "./templates")

	indexHTML, err := box.FindString("index.html")
	if err != nil {
		e.Logger.Fatal(err)
	}

	t := template.Must(template.New("index.html").Parse(indexHTML))

	r := &Template{templates: t}
	e.Renderer = r

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/main.js", func(c echo.Context) error {
		js, err := box.Find("main.js")
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "application/javascript", js)
	})
	e.GET("/styles.css", func(c echo.Context) error {
		css, err := box.Find("styles.css")
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "text/css", css)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
