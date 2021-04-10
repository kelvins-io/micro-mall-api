package goroutine

import (
	"log"
	"runtime/debug"
)

func AttachPanicHandle(f func()) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("goroutine panic: %v, stacktrace:%v", err, string(debug.Stack()))
			}
		}()
		f()
	}
}
