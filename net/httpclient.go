package net

import (
	"bytes"
	"github.com/hujun8610/gomylibrary/logger"
	"io"
	"net/http"
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
		log.Fatalf("get response failed", err)
	}

	log.WithField("StatusCode", resp.StatusCode).Info("http status code")
	defer resp.Body.Close()
	log.WithField("responseHeader", resp.Header).Info("response headers")

	data, _ := io.ReadAll(resp.Body)
	log.WithField("responseBody", string(data)).Info("response")
	return string(data)
}
