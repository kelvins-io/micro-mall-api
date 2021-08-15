package startup

import "runtime"

func execProcessCmd(pid int, upType startUpType) (next bool, err error) {
	cmd := *control
	switch cmd {
	case startUpReStart:
		logging.Infof("process platform(%s) not support restart\n", runtime.GOOS)
	case startUpStop:
		logging.Info("process stop...")
		err = processControl(pid, syscall.SIGTERM)
		logging.Info("process stop over")
	default:
		next = true
	}
	return
}
