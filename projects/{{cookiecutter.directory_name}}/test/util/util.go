package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

type ExpectedAndGot struct {
	PropertyName string
	Expects      []interface{}
	Got          []interface{}
}

func NotExpected(expectedAndGots []ExpectedAndGot) string {
	result := "Got unexpected result \n"
	template := "Expected property %s = \"%v\" but got %s = \"%v\" \n"

	for _, expectedAndGot := range expectedAndGots {
		for index, expected := range expectedAndGot.Expects {
			if len(expectedAndGot.Got) > index {
				result += fmt.Sprintf(template, expectedAndGot.PropertyName, expected, expectedAndGot.PropertyName, expectedAndGot.Got[index])
			} else {
				result += fmt.Sprintf(template, expectedAndGot.PropertyName, expected, expectedAndGot.PropertyName, "")
			}
		}
	}

	return result
}

func CommonRequest(requestUrl, httpMethod string, postBody json.RawMessage, header map[string]string, skipTlsCheck, disableKeepAlive bool, timeout time.Duration) ([]byte, int, error) {
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
	client.Timeout = timeout
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

func SameStructIsExpected(t *testing.T, expected, got interface{}) bool {
	var failedList []ExpectedAndGot
	valueSetsA := reflect.ValueOf(expected).Elem()
	valueSetsB := reflect.ValueOf(got).Elem()
	for i := 0; i < valueSetsA.NumField(); i++ {
		valueField := valueSetsA.Field(i)
		typeField := valueSetsA.Type().Field(i)
		if !reflect.DeepEqual(valueField.Interface(), valueSetsB.Field(i).Interface()) {
			failedList = append(failedList, ExpectedAndGot{
				PropertyName: typeField.Name,
				Expects:      []interface{}{valueField.Interface()},
				Got:          []interface{}{valueSetsB.Field(i).Interface()},
			})
		}
	}
	if len(failedList) > 0 {
		t.Error(NotExpected(failedList))
	}
	return len(failedList) == 0
}

func HttpStatusIsExpected(t *testing.T, expectedStatus []int, got int) bool {
	ok := false
	for _, eStatus := range expectedStatus {
		if eStatus == got {
			ok = true
			break
		}
	}
	if !ok {
		t.Error(NotExpected([]ExpectedAndGot{
			{
				PropertyName: "http status",
				Expects:      []interface{}{expectedStatus},
				Got:          []interface{}{got},
			},
		}))
	}
	return ok
}

func IntIsExpected(t *testing.T, propertyName string, expectedIntegers []int, gots []int) bool {
	ok := false
	for _, expectedInteger := range expectedIntegers {
		for _, got := range gots {
			if expectedInteger == got {
				ok = true
				break
			}
		}
	}
	if !ok {
		t.Error(NotExpected([]ExpectedAndGot{
			{
				PropertyName: propertyName,
				Expects:      []interface{}{expectedIntegers},
				Got:          []interface{}{gots},
			},
		}))
	}
	return ok
}

func IntGreateThanOrGreateThanEqual(t *testing.T, propertyName string, expectedGreateThan, got int, equal bool) bool {
	template := "Expected property %s value %s %v but got %v"
	if equal {
		if got < expectedGreateThan {
			t.Error(fmt.Sprintf(template, propertyName, ">=", expectedGreateThan, got))
			return false
		}
	} else {
		if got <= expectedGreateThan {
			t.Error(fmt.Sprintf(template, propertyName, ">", expectedGreateThan, got))
			return false
		}
	}
	return true
}

func IntLessThanOrLessThanEqual(t *testing.T, propertyName string, expectedGreateThan, got int, equal bool) bool {
	template := "Expected property %s value %s %v but got %v"
	if equal {
		if got > expectedGreateThan {
			t.Error(fmt.Sprintf(template, propertyName, "<=", expectedGreateThan, got))
			return false
		}
	} else {
		if got >= expectedGreateThan {
			t.Error(fmt.Sprintf(template, propertyName, "<", expectedGreateThan, got))
			return false
		}
	}
	return true
}

func StringExpected(t *testing.T, propertyName string, expectedStrings []string, gots []string) bool {
	ok := false
	for _, expectedString := range expectedStrings {
		for _, got := range gots {
			if expectedString == got {
				ok = true
				break
			}
		}
	}
	if !ok {
		t.Error(NotExpected([]ExpectedAndGot{
			{
				PropertyName: propertyName,
				Expects:      []interface{}{expectedStrings},
				Got:          []interface{}{gots},
			},
		}))
	}
	return ok
}

func TryJsonUnMarsh(t *testing.T, data []byte, value interface{}, structName string) bool {
	err := json.Unmarshal(data, value)
	if err != nil {
		t.Errorf("Failed to parse json to struct %s due to %v", structName, err)
		return false
	}
	return true
}
