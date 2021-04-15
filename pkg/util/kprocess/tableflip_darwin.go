// +build amd64,darwin

package kprocess

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudflare/tableflip"
)

type KProcess struct {
	pidFile string
	processUp *tableflip.Upgrader
}

// This shows how to use the upgrader
// with the graceful shutdown facilities of net/http.
func (k *KProcess) Listen(network, addr, pidFile string) (ln net.Listener, err error) {
	log.Printf(fmt.Sprintf("[info] exec process pid %d \n", os.Getpid()))

	k.processUp, err = tableflip.New(tableflip.Options{
		UpgradeTimeout: 500 * time.Millisecond,
		PIDFile:        pidFile,
	})
	if err != nil {
		return nil, err
	}
	k.pidFile = pidFile

	go k.signal(k.upgrade, k.stop)

	// Listen must be called before Ready
	if network != "" && addr != "" {
		ln, err = k.processUp.Listen(network, addr)
		if err != nil {
			return nil, err
		}
	}
	if err := k.processUp.Ready(); err != nil {
		return nil, err
	}

	return ln, nil
}

func (k *KProcess) stop() error {
	if k.processUp != nil {
		k.processUp.Stop()
		return os.Remove(k.pidFile)
	}
	return nil
}

func (k *KProcess) upgrade() error {
	if k.processUp != nil {
		return k.processUp.Upgrade()
	}
	return nil
}

func (k *KProcess) Exit() <-chan struct{} {
	if k.processUp != nil {
		return k.processUp.Exit()
	}
	ch := make(chan struct{})
	close(ch)
	return ch
}

func (k *KProcess) signal(upgradeFunc, stopFunc func() error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	for s := range sig {
		switch s {
		case syscall.SIGTERM:
			if stopFunc != nil {
				err := stopFunc()
				if err != nil {
					log.Println("KProcessUp stopFunc failed:", err)
				}
			}
			return
		case syscall.SIGUSR1, syscall.SIGUSR2:
			if upgradeFunc != nil {
				err := upgradeFunc()
				if err != nil {
					log.Println("KProcessUp Upgrade failed:", err)
				}
			}
		}
	}
}
