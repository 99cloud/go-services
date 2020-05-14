package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/logger"
)

func Post(url string, data interface{}) (int, string) {
	timeout := 5 * time.Second
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true},
	}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		logger.Error(nil, err.Error())
		return 0, ""
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(result)
}
