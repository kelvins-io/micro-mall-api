package startup

import (
	"flag"
	"gitee.com/cristiane/micro-mall-api/internal/logging"
	"io/ioutil"
	"os"
	"strconv"
)

type startUpType string

const (
	startUpStart   startUpType = "start"
	startUpReStart startUpType = "restart"
	startUpStop    startUpType = "stop"
)

var (
	control = flag.String("s", string(startUpStart), "control cmd eg: start，stop，restart")
)

func processControl(pid int, signal os.Signal) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Signal(signal)
}

func ParseCliCommand(pidFile string) (next bool, err error) {
	flag.Parse()
	cmd := startUpType(*control)
	switch cmd {
	case startUpStart:
		next = true
		return
	case startUpReStart:
	case startUpStop:
	default:
		next = false
		logging.Info("unsupported command!!!")
		return
	}

	pid, err := parsePidFile(pidFile)
	if err != nil {
		return
	}

	return execProcessCmd(pid, cmd)
}

func parsePidFile(pidFile string) (pid int, err error) {
	_, err = os.Stat(pidFile)
	if err != nil {
		return
	}
	var f *os.File
	f, err = os.OpenFile(pidFile, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	pid, err = strconv.Atoi(string(content))
	return
}
