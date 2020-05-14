package v1

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/internal/APP_46ea591951824d8e9376b0f98fe4d48a/controller/resources"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/internal/APP_46ea591951824d8e9376b0f98fe4d48a/model"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/constants"
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
				model.CommonResponse{},
			},
		},
		Tags: []string{constants.TAG_COMMON_echo},
	},
}
