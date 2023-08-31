package qimen

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"sort"
)

const (
	SIGN_METHOD_MD5  = "md5"
	SIGN_METHOD_HMAC = "hmac"
)

func SignTopRequest(params map[string]string, body string, secret string, signMethod string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var query bytes.Buffer
	if signMethod == SIGN_METHOD_MD5 {
		query.WriteString(secret)
	}
	for _, key := range keys {
		value := params[key]
		if key != "" && value != "" {
			query.WriteString(key)
			query.WriteString(value)
		}
	}

	if body != "" {
		query.WriteString(body)
	}

	var bytes []byte
	if signMethod == SIGN_METHOD_HMAC {
		bytes = encryptHMAC(query.String(), secret)
	} else {
		query.WriteString(secret)
		bytes, _ = encryptMD5(query.Bytes())
	}

	return byte2hex(bytes)
}

func encryptHMAC(data string, secret string) []byte {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

func encryptMD5(data []byte) ([]byte, error) {
	hash := md5.New()
	_, err := hash.Write(data)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func byte2hex(bytes []byte) string {
	sign := ""
	for _, b := range bytes {
		hex := fmt.Sprintf("%02X", b)
		sign += hex
	}
	return sign
}
