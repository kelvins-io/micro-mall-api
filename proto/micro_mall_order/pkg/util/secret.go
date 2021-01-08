package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
)

// 生成秘钥
func GenerateSecret() string {
	m := md5.New()
	m.Write(uuid.NewV4().Bytes())
	return fmt.Sprintf("%x", m.Sum(nil))
}

func GenerateSha1Sign(params, appSecret string) string {
	s := fmt.Sprintf("%v%v", params, appSecret)
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
