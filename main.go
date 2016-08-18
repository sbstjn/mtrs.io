package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/marpaia/graphite-golang"
)

func main() {
	port := os.Getenv("PORT")
	graphiteKey := os.Getenv("HOSTEDGRAPHITE_APIKEY")
	Graphite, _ := graphite.NewGraphite("df6c1ae2.carbon.hostedgraphite.com", 2003)
	Graphite.SimpleSend(graphiteKey+".stats.graphite_loaded", "1")

	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		log.Print("Using port ", port)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/_/:factor/:rating", func(c *gin.Context) {
		factor := c.Param("factor")
		rating := c.Param("rating")

		widthFactor := 4 + len(factor)*8
		widthRating := 4 + len(rating)*8

		c.HTML(http.StatusOK, "shield.svg", gin.H{
			"factor":         factor,
			"rating":         rating,
			"width":          widthFactor + widthRating,
			"widthFactor":    widthFactor,
			"widthRating":    widthRating,
			"positionFactor": widthFactor / 2,
			"positionRating": widthFactor + (widthRating / 2),
		})
	})

	router.Run(":" + port)
	/*
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
	*/
}
