package v1

import (
	"{{cookiecutter.project_slug}}/internal/{{cookiecutter.app_slug}}/controller/resources"
	"{{cookiecutter.project_slug}}/pkg/constants"
	"{{cookiecutter.project_slug}}/pkg/schema"
	"encoding/json"
	"net/http"

	restfulSpec "github.com/emicklei/go-restful-openapi"
)

var echo = []Route{
	{
		Path:           "/echo",
		HTTPMethod:     http.MethodPost,
		Handler:        resources.Echo,
		ChallengeCode:  255,
		Doc:            "This method is used to echo.",
		PathParams:     nil,
		MetaData:       restfulSpec.KeyOpenAPITags,
		ReadDataModel:  json.RawMessage{},
		WriteDataModel: json.RawMessage{},
		ReturnDefinitions: []DocReturnDefinition{
			{
				http.StatusOK,
				constants.HTTP_200,
				json.RawMessage{},
			},
			{
				http.StatusBadRequest,
				constants.HTTP_400,
				schema.CommonResponse{},
			},
		},
		Tags: []string{constants.TAG_COMMON_echo},
	},
}
