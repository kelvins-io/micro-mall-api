package metrics_mux

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func GetPrometheusMux(mux *http.ServeMux) *http.ServeMux {
	mux.Handle("/metrics", promhttp.Handler())
	return mux
}
