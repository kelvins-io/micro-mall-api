package goroutine

import "context"

func CheckGoroutineErr(errCtx context.Context) error {
	select {
	case <-errCtx.Done():
		return errCtx.Err()
	default:
		return nil
	}
}
