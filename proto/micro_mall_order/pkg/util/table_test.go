package util

import (
	"testing"
)

func TestSeparateTable(t *testing.T) {
	var id string
	var count int
	if SeparateTable(id, count) != "" {
		t.Error("空字符串返回值不为空")
	}

	id = "a"
	if SeparateTable(id, count) != "" {
		t.Error("分表数为0时返回值不为空")
	}

	count = 1
	if SeparateTable(id, count) == "" {
		t.Error("返回值为空")
	}

	id = "adasd44---中"
	count = 128
	if SeparateTable(id, count) == "" {
		t.Error("测试特殊值")
	}
}
