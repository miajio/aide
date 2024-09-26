package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

// md5
func Md5(content string) (md string) {
	if content != "" {
		h := md5.New()
		_, _ = io.WriteString(h, content)
		md = fmt.Sprintf("%x", h.Sum(nil))
	}
	return
}
