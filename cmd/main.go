package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	precompiler "github.com/parnic/go-assetprecompiler"
)

func main() {

	precompiler.Compile(precompiler.Config{
		Files: []string{
			"static/css/bootstrap.min.css",
			"static/sticky-footer-navbar.css",
			"static/base-custom.css",
		},
		Minify:    true,
		OutputDir: "static/",
	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("cmd/templates/*.tmpl.html")
	router.Static("/static", "cmd/static")
	router.Use(precompiler.GinMiddleware("/static"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/sources", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sources.tmpl.html", nil)
	})

	router.GET("/destinations", func(c *gin.Context) {
		c.HTML(http.StatusOK, "destinations.tmpl.html", nil)
	})

	router.Run(":" + port)
}
