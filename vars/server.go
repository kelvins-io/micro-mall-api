package vars

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Version    = "1.0.0"
	AppTypeWeb = 1
)

// Application ...
type Application struct {
	Name       string
	Type       int32
	LoadConfig func() error
	SetupVars  func() error
	StopFunc   func() error
}

// ListenerApplication ...
type WEBApplication struct {
	*Application
	EndPort        int
	MonitorEndPort int
	// 监控
	Mux *http.ServeMux
	// RegisterHttpRoute 定义HTTP router
	RegisterHttpRoute func() *gin.Engine
	// 系统定时任务
	RegisterTasks func() []CronTask
}

type ListenerApplication struct {
	*Application
	EndPort        int
	MonitorEndPort int
	Network        string
	ReadTimeout    int
	WriteTimeout   int
	// 监控
	Mux *http.ServeMux
	// Listener Server Accept 后的自定义事件
	EventHandler func(context.Context, []byte) ([]byte, error)
}
