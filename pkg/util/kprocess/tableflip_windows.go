// +build windows

package kprocess

import (
	"fmt"
	"gitee.com/cristiane/micro-mall-api/internal/logging"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type KProcess struct {
	pidFile string
	pid     int
	ch      chan struct{}
}

// This shows how to use the upgrader
// with the graceful shutdown facilities of net/http.
func (k *KProcess) Listen(network, addr, pidFile string) (ln net.Listener, err error) {
	if k.ch == nil {
		k.ch = make(chan struct{})
	}
	k.pid = os.Getpid()
	logging.Info(fmt.Sprintf("exec process pid %d \n", k.pid))
	logging.Info("warning windows only support process shutdown ")

	go k.signal(k.stop)

	return net.Listen(network, addr)
}

func (k *KProcess) stop() error {
	close(k.ch)
	return nil
}

func (k *KProcess) upgrade() error {
	return nil
}

func (k *KProcess) Exit() <-chan struct{} {
	return k.ch
}

func (k *KProcess) signal(stopFunc func() error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)
	for s := range sig {
		switch s {
		case syscall.SIGTERM:
			if stopFunc != nil {
				err := stopFunc()
				if err != nil {
					logging.Infof("KProcess exec stopFunc failed:%v\n", err)
				}
				logging.Infof("process [%d] stop...\n", k.pid)
			}
			return
		}
	}
}
