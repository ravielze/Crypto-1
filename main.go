package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/affine"
	"github.com/ravielze/Crypto-1/internal/hill"
	"github.com/ravielze/Crypto-1/internal/playfair"
	"github.com/ravielze/Crypto-1/internal/vignere"
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

	e.GET("/starter-template.css", func(c echo.Context) error {
		css, err := box.Find("starter-template.css")
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "text/css", css)
	})

	e.POST("/api/alphabet/:name/:type", func(c echo.Context) error {
		type Request struct {
			Text string `json:"text"`
			Key  string `json:"key"`
			M    int    `json:"m"`
			N    int    `json:"n"`
		}
		var req Request
		if err := c.Bind(&req); err != nil {
			return err
		}
		fmt.Println(req)
		if len(req.Key) == 0 && c.Param("name") != "affine" {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "key must not empty"})
		}
		method := map[string]internal.AlphabetCipher{
			"stdvignere": vignere.NewStandard(req.Key),
			"akyvignere": vignere.NewAutoKey(req.Key),
			"extvignere": vignere.NewExtendedString(req.Key),
			"affine":     affine.NewString(req.M, req.N),
			"playfair":   playfair.NewPlayfair(req.Key),
			"hill":       hill.NewHill(req.Key),
		}
		usedMethod, ok := method[c.Param("name")]
		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "method name not found"})
		}
		var result string

		if c.Param("type") == "encrypt" {
			result = usedMethod.Encrypt(req.Text)
		} else {
			result = usedMethod.Decrypt(req.Text)
		}

		if c.Param("name") == "affine" && len(result) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": "m and 256 are not coprime."})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"result": result,
		})
	})

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":8000"))
}
