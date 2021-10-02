package middleware

import (
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

func RateLimit(maxConcurrent int) gin.HandlerFunc {
	var limiter Limiter
	if maxConcurrent > 0 {
		limiter = NewKelvinsRateLimit(maxConcurrent)
	}
	return func(c *gin.Context) {
		if limiter != nil {
			if limiter.Limit() {
				app.JsonResponse(c, http.StatusTooManyRequests, code.TooManyRequests, code.GetMsg(code.TooManyRequests))
				c.Abort()
				return
			}
			defer limiter.ReturnTicket()
		}
		c.Next()
	}
}

func NewKelvinsRateLimit(maxConcurrent int) Limiter {
	limiter := &kelvinsRateLimit{}
	if maxConcurrent > 0 {
		limiter.maxConcurrent = maxConcurrent
		limiter.tickets = make(chan struct{}, maxConcurrent+1)
	}
	return limiter
}

type kelvinsRateLimit struct {
	maxConcurrent int
	tickets       chan struct{}
	ticketsState  int32
}

func (r *kelvinsRateLimit) Limit() bool {
	// no limit
	if r.maxConcurrent == 0 {
		return false
	}
	// take ticket
	take := r.takeTicket()
	if take {
		return false
	}

	return true
}

func (r *kelvinsRateLimit) takeTicket() bool {
	if r.maxConcurrent == 0 {
		return true
	}
	if r.tickets == nil {
		return true
	} else {
		if atomic.LoadInt32(&r.ticketsState) == 1 {
			return false
		}
	}

	select {
	case r.tickets <- struct{}{}:
		return true
	case <-vars.AppCloseCh:
		atomic.StoreInt32(&r.ticketsState, 1)
		close(r.tickets)
		return false
	default:
		return false
	}
}

func (r *kelvinsRateLimit) ReturnTicket() {
	if r.maxConcurrent == 0 {
		return
	}
	if r.tickets == nil {
		return
	}
	select {
	case <-r.tickets:
	default:
	}
}

type Limiter interface {
	Limit() bool
	ReturnTicket()
}
