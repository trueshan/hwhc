package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func Sh1(AppSer, Nonce, Timestamp string) string {
	h := sha1.New()
	h.Write([]byte(AppSer + Nonce + Timestamp))
	ctsignature := fmt.Sprintf("%x", h.Sum(nil))
	return ctsignature

}

func GetHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
