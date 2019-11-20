package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	job "github.com/Radabaugh/funnel/lib/job"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("FUNNEL_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env)
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("cmd/templates/*.tmpl.html")
	router.Static("/static", "cmd/static")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/jobs/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new-job-form.tmpl.html", nil)
	})

	router.POST("/", func(c *gin.Context) {

		source := c.PostForm("job[source]")
		destination := c.PostForm("job[destination]")
		interval := c.PostForm("job[interval]")
		frequency, err := strconv.Atoi(c.PostForm("job[frequency]"))
		if err != nil {
			panic(err)
		}

		job.Save(source, destination, interval, frequency)
	})

	router.GET("/sources", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sources.tmpl.html", nil)
	})

	router.GET("/sources/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new-source-form.tmpl.html", nil)
	})

	router.GET("/destinations", func(c *gin.Context) {
		c.HTML(http.StatusOK, "destinations.tmpl.html", nil)
	})

	router.GET("/destinations/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new-destination-form.tmpl.html", nil)
	})

	router.Run(":" + port)
}
