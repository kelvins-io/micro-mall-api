package startup

func execProcessCmd(pid int, upType startUpType) (next bool, err error) {
	switch upType {
	case startUpReStart:
		logging.Info("process restart...")
		err = processControl(pid, syscall.SIGUSR1)
		logging.Info("process restart over")
	case startUpStop:
		logging.Info("process stop...")
		err = processControl(pid, syscall.SIGTERM)
		logging.Info("process stop over")
	default:
		next = true
	}
	return
}
