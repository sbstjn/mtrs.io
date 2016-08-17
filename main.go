package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/shield/:name/:value", func(c *gin.Context) {
		name := c.Param("name")
		value := c.Param("value")

		c.HTML(http.StatusOK, "shield.svg", gin.H{
			"name":  name,
			"value": value,
		})
	})

	router.GET("/rewrite/:name/:value", func(c *gin.Context) {
		name := c.Param("name")
		value := c.Param("value")

		c.Redirect(http.StatusMovedPermanently, "https://img.shields.io/badge/"+name+"-"+value+"-FFA800.svg")
	})

	router.Run(":" + port)

}
