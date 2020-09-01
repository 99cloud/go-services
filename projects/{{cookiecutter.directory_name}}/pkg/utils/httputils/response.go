package httputils

import (
	"{{cookiecutter.project_slug}}/pkg/constants"
	"{{cookiecutter.project_slug}}/pkg/logger"
	"{{cookiecutter.project_slug}}/pkg/schema"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
)

func WriteCommonResponse(resp *restful.Response, status int, message, apiName string) {
	logger.Debug(nil, message)

	body := schema.CommonResponse{
		Type:     "Common",
		Title:    "Error",
		Status:   status,
		Detail:   message,
		Instance: apiName,
	}

	resp.WriteHeaderAndJson(status, body, restful.MIME_JSON)
}

func WriteCommonInternalError(resp *restful.Response, err error, apiName string) {
	logger.Error(nil, fmt.Sprintf(err.Error()))
	WriteCommonResponse(resp, http.StatusInternalServerError, constants.ERR_MSG_INTERNAL_ERR, apiName)
}

func WriteStatusAndEntityWithLog(resp *restful.Response, httpStatus int, value interface{}, err error) {
	logger.Error(nil, fmt.Sprintf(err.Error()))
	_ = resp.WriteHeaderAndEntity(httpStatus, value)
}

func WriteSingleEmptyStructWithStatus(resp *restful.Response, status int) {
	resp.WriteHeader(status)
	_, _ = resp.Write([]byte("{}"))
}

func WriteEmptyListWithStatus(resp *restful.Response, status int) {
	resp.WriteHeader(status)
	_, _ = resp.Write([]byte("[]"))
}
