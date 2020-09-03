package setup

import (
	"gitee.com/cristiane/micro-mall-api/internal/metrics_mux"
	"net/http"
)

// NewServerMux ...
func NewServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux = metrics_mux.GetElasticMux(mux)
	mux = metrics_mux.GetPProfMux(mux)
	mux = metrics_mux.GetPrometheusMux(mux)
	return mux
}
