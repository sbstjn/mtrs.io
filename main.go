package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "mtrs",
			Name:      "request",
			Help:      "Total number of requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		log.Print("Using port ", port)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Done.")
		requestCounter.WithLabelValues("/").Inc()
	})

	r.Group("/metrics", gin.BasicAuth(gin.Accounts{
		"prometheus": "secret",
	})).GET("/", func(c *gin.Context) {
		prometheus.Handler().ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":" + port)
}
