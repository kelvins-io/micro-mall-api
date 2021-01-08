package util

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const InternalCode = "86"

func GetCompletePhone(countryCode, mobile string) string {
	if mobile == "" {
		return ""
	}

	if countryCode == InternalCode || countryCode == "" {
		return mobile
	}

	return fmt.Sprintf("00-%v-%v", countryCode, mobile)
}

func GenerateSignCode(count uint8) string {
	if count == 0 {
		return ""
	}
	var buffer bytes.Buffer
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := uint8(0); i < count; i++ {
		buffer.WriteString(strconv.Itoa(rnd.Intn(10)))
	}
	return buffer.String()
}
