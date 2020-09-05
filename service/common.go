package service

import (
	"github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
	"strconv"
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
