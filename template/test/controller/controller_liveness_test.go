package controller

import (
	"net/http"
	"testing"

	testUtil "PROJECT_46ea591951824d8e9376b0f98fe4d48a/test/util"
)

func TestControllerLiveness(t *testing.T) {
	// wait 3 seconds for the server to starts
	//time.Sleep(time.Second * 3)
	t.Run("Got echo from liveness", testLivenessEcho)
}

func testLivenessEcho(t *testing.T) {
	var postDataEcho = `{
		"test": {
			"test1": {
				"test2": 43
			},
			"xyz": true
		},
		"test3": "test4"
	}`
	requestURL := "http://localhost:8080/liveness/v1/echo"
	jsonRaw := []byte(postDataEcho)
	_, httpStatus := requestTest(t, requestURL, http.MethodPost, jsonRaw, map[string]string{}, true, true)
	testUtil.HttpStatusIsExpected(t, []int{http.StatusOK}, httpStatus)
}
