package main

import (
	"net/http"

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
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Done.")
		requestCounter.WithLabelValues("/").Inc()
	})

	authorized := r.Group("/metrics", gin.BasicAuth(gin.Accounts{
		"prometheus": "secret",
	}))

	authorized.GET("/", func(c *gin.Context) {
		prometheus.Handler().ServeHTTP(c.Writer, c.Request)
	})

	r.Run("localhost:5000")
}
