package buf

import "strings"

func GenerateBuf() *strings.Builder {
	buf := &strings.Builder{}
	return buf
}

func ResetBuf(buf *strings.Builder) {
	buf.Reset()
}

func WriteString(buf *strings.Builder, params ...string) {
	for _, str := range params {
		buf.WriteString(str)
	}
}

func ReadBuf(buf *strings.Builder) string {
	return buf.String()
}
