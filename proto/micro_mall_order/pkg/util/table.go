package util

import (
	"fmt"
	"strconv"
)

// 分表
func SeparateTable(id string, count int) string {
	if id == "" || count <= 0 {
		return ""
	}

	idInt := 0
	for _, value := range []byte(id) {
		idInt += int(value)
	}

	length := len(strconv.Itoa(count))
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"s", strconv.Itoa(idInt%count))
}
