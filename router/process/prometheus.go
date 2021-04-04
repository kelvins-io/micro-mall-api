package process

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsApi(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}
