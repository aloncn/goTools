package gotool

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"time"
)

var F = com{}
type com struct {

}
func (c *com) Hi() string {
	return "hi"
}

// 判断邮箱
func (c *com) IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 判断手机号
func (c *com) IsMobile(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5-7,9])|(15[0-3,5-9])|(17[0-3,5-8])|(18[0-9])|(19[0-9])|166)\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func (c *com) RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (c *com) MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}