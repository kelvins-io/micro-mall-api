package service

import (
	"strconv"

	"gitee.com/cristiane/micro-mall-api/pkg/code"
	uuid "github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
)

func genUUID() string {
	return uuid.NewV4().String()
}

func GenInviterCode() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return genUUID()
	}
	return strconv.FormatUint(id, 16)
}

func GenAccountId() string {
	return genUUID()
}

func errorToRetCode(err error) (retCode int) {
	if err == nil {
		retCode = code.SUCCESS
		return
	}
	errRet := err.Error()
	if errRet == "" {
		retCode = code.ERROR
		return
	}
	retCode, errAoi := strconv.Atoi(errRet)
	if errAoi != nil {
		retCode = code.ERROR
		return
	}
	return
}
