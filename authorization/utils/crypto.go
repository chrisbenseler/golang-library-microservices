package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMd5 get Md5 from string
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
