package metrics_mux

import (
	"expvar"
	"fmt"
	"gitee.com/kelvins-io/common/ptool"
	"net/http"
)

var appStats = expvar.NewMap("appstats")

func GetElasticMux(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/debug/vars", metricsHandler)
	return mux
}

// metricsHandler print expvar data.
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	appStats.Set("Goroutine", expvar.Func(ptool.GetGoroutineCount))
	appStats.Set("Threadcreate", expvar.Func(ptool.GetThreadCreateCount))
	appStats.Set("Block", expvar.Func(ptool.GetBlockCount))
	appStats.Set("Mutex", expvar.Func(ptool.GetMutexCount))
	appStats.Set("Heap", expvar.Func(ptool.GetHeapCount))

	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(w, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(w, "%q: %q", key, str)
		} else {
			fmt.Fprintf(w, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(w, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(w, "\n}\n")
}
