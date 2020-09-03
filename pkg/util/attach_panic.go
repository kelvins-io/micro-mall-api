package util

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/vars"
	"runtime/debug"
)

func AttachPanicHandle(f func()) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				vars.ErrorLogger.Errorf(context.Background(), "goroutine panic: %v, stacktrace:%v", err, string(debug.Stack()))
			}
		}()
		f()
	}
}
