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

var (
	processUp *tableflip.Upgrader
)

// This shows how to use the upgrader
// with the graceful shutdown facilities of net/http.
func Listen(network, addr, pidFile string) (ln net.Listener, err error) {
	log.Printf(fmt.Sprintf("exec process pid %d \n", os.Getpid()))

	processUp, err = tableflip.New(tableflip.Options{
		UpgradeTimeout: 500 * time.Millisecond,
		PIDFile:        pidFile,
	})
	if err != nil {
		return nil, err
	}

	go Signal(Upgrade, Stop)

	// Listen must be called before Ready
	if network != "" && addr != "" {
		ln, err = processUp.Listen(network, addr)
		if err != nil {
			return nil, err
		}
	}
	if err := processUp.Ready(); err != nil {
		return nil, err
	}

	return ln, nil
}

func Stop() error {
	if processUp != nil {
		processUp.Stop()
	}
	return nil
}

func Upgrade() error {
	if processUp != nil {
		return processUp.Upgrade()
	}
	return nil
}

func Exit() <-chan struct{} {
	return processUp.Exit()
}

func Signal(upgradeFunc, stopFunc func() error) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM)
	for s := range sig {
		switch s {
		case syscall.SIGTERM:
			if stopFunc != nil {
				err := stopFunc()
				if err != nil {
					log.Println("ProcessUp stopFunc failed:", err)
				}
			}
			return
		case syscall.SIGUSR1, syscall.SIGUSR2:
			if upgradeFunc != nil {
				err := upgradeFunc()
				if err != nil {
					log.Println("ProcessUp Upgrade failed:", err)
				}
			}
		}
	}
}
