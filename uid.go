package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

// convertUID は uid を隠ぺいするため一方向ハッシュを行います
func convertUID(r *http.Request, uid string) string {
	var (
		salt   = []byte("stringnote")
		strech = 13
	)
	if conf, err := getConfig(r); err == nil {
		salt = []byte(conf.Salt)
		strech = conf.Strech
	}

	convined := bytes.Join([][]byte{salt, []byte(uid)}, []byte{})
	sum := sha256.Sum256(convined)
	for i := 0; i < strech-1; i++ {
		sum = sha256.Sum256(sum[0:32])
	}
	// 長いので短縮
	sum1 := sha1.Sum(sum[0:32])
	ret := strings.Trim(base64.StdEncoding.EncodeToString(sum1[0:16]), "=")
	ret = strings.Replace(ret, "+", "-", -1)
	ret = strings.Replace(ret, "/", "_", -1)
	return ret
}
