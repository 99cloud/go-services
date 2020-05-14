package v1

import (
	"github.com/emicklei/go-restful"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/schema"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/runtime"
)

type Path struct {
	GroupVersion runtime.GroupVersion
	Routes       []Route
}

type Route struct {
	Path              string
	HTTPMethod        string
	Handler           func(request *restful.Request, response *restful.Response)
	ChallengeCode     int
	Doc               string
	PathParams        []schema.Parameter
	QueryParams       []schema.Parameter
	Filter            []restful.FilterFunction
	MetaData          string
	ReadDataModel     interface{}
	WriteDataModel    interface{}
	ReturnDefinitions []DocReturnDefinition
	Tags              []string
}

type DocReturnDefinition struct {
	HTTPStatus  int
	Message     string
	ReturnModel interface{}
}

var paths = []Path{
	{
		GroupVersion: runtime.GroupVersion{Group: "liveness", Version: "v1"},
		Routes:       echo,
	},
}
