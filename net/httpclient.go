package net

import (
	"bytes"
	"github.com/hujun8610/gomylibrary/logger"
	"io"
	"net/http"
	"strings"
)

var log = logger.GetLogger()

func PostRequest(url string, headers map[string]string, payload string) string {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Fatalf("request failed")
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("get response failed %v", err)
	}

	log.WithField("StatusCode", resp.StatusCode).Info("http status code")
	defer resp.Body.Close()
	log.WithField("responseHeader", resp.Header).Info("response headers")

	data, _ := io.ReadAll(resp.Body)
	log.WithField("responseBody", string(data)).Info("response")
	return string(data)
}

func PostFormRequest(urlstr string) string {
	str := strings.Split(urlstr, "?")
	if len(str) != 2 {
		log.Errorf("split url string failed %v", str)
		panic(nil)
	}
	return PostFormRequestWithPayLoad(str[0], str[1])
}

func PostFormRequestWithPayLoad(baseUrl string, payLoad string) string {
	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBufferString(payLoad))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("post request failed %v", err)
		panic(err)
	}
	defer resp.Body.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		log.Errorf("read data from resp body failed")
		panic(err)
	}

	return buffer.String()
}
