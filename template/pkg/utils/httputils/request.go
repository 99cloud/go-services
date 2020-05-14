package httputils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func CommonRequest(requestUrl, httpMethod string, postBody json.RawMessage, header map[string]string, skipTlsCheck, disableKeepAlive bool) ([]byte, int, error) {
	var req *http.Request
	var reqErr error

	req, reqErr = http.NewRequest(httpMethod, requestUrl, bytes.NewReader(postBody))

	if reqErr != nil {
		return []byte{}, http.StatusInternalServerError, reqErr
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for key, val := range header {
		req.Header.Add(key, val)
	}
	client := &http.Client{}
	client.Timeout = 5 * time.Second
	if skipTlsCheck {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, DisableKeepAlives: disableKeepAlive}
	} else {
		client.Transport = &http.Transport{DisableKeepAlives: disableKeepAlive}
	}
	resp, respErr := client.Do(req)
	if respErr != nil {
		return []byte{}, http.StatusInternalServerError, respErr
	}
	defer resp.Body.Close()
	body, readBodyErr := ioutil.ReadAll(resp.Body)
	if readBodyErr != nil {
		return []byte{}, http.StatusInternalServerError, readBodyErr
	}
	return body, resp.StatusCode, nil
}
