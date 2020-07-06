package resources

import (
	"encoding/json"
	"net/http"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/constants"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/schema"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/utils/httputils"

	"github.com/emicklei/go-restful"
)

var MOCK_QUEUE chan MockResult

type MockResult struct {
	Result bool
}

func Echo(req *restful.Request, resp *restful.Response) {
	body := json.RawMessage{}
	if err := req.ReadEntity(&body); err != nil {
		httputils.WriteCommonResponse(resp, http.StatusBadRequest, err.Error(), constants.TAG_COMMON_echo)
		return
	}
	resp.WriteAsJson(body)
}

func MockError(req *restful.Request, resp *restful.Response) {
	resp.WriteHeader(http.StatusInternalServerError)
}

func MockPostBack(req *restful.Request, resp *restful.Response, result bool) *MockResult {
	body := schema.CommonResponse{}
	if err := req.ReadEntity(&body); err != nil {
		httputils.WriteCommonResponse(resp, http.StatusBadRequest, err.Error(), constants.TAG_COMMON_mock)
		return nil
	}

	resp.WriteHeaderAndJson(http.StatusOK, body, restful.MIME_JSON)

	r := MockResult{result}
	return &r
}

func MockSuccess(req *restful.Request, resp *restful.Response) {
	// Add req to cache queue
	r := MockPostBack(req, resp, true)

	go func() {
		MOCK_QUEUE <- *r
	}()
}

func MockFailure(req *restful.Request, resp *restful.Response) {
	// Add req to cache queue
	r := MockPostBack(req, resp, false)

	go func() {
		MOCK_QUEUE <- *r
	}()
}

func init() {
	MOCK_QUEUE = make(chan MockResult, 10)
}
