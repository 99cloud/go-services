package v1

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/params"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/runtime"

	"github.com/emicklei/go-restful"
)

//const GroupName = ""
const GroupName = "resources.io"

//var GroupVersion = runtime.GroupVersion{Group: GroupName, Version: ""}
//var GroupVersion = runtime.GroupVersion{Group: GroupName, Version: "v1"}

var (
	WebServiceBuilder = runtime.NewContainerBuilder(addWebServiceWithStaticRoute)
	AddToContainer    = WebServiceBuilder.AddToContainer
)

//func addWebService(c *restful.Container) error {
//	webservice := runtime.NewWebService(GroupVersion)
//
//	ok := "ok"
//
//	webservice.Route(webservice.GET("/testrestful/{resources}").
//		To(resources.TestRestful).
//		Metadata(restfulspec.KeyOpenAPITags, []string{constants.TestResourcesTag}).
//		Doc("test restful query").
//		Param(webservice.PathParameter("resources", "namespace level resource type, e.g. pods,jobs,configmaps,services.")).
//		Param(webservice.QueryParameter(params.ConditionsParam, "query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a").
//			Required(false).
//			DataFormat("key=%s,key~%s")).
//		Param(webservice.QueryParameter(params.PagingParam, "paging query, e.g. limit=100,page=1").
//			Required(false).
//			DataFormat("limit=%d,page=%d").
//			DefaultValue("limit=10,page=1")).
//		Param(webservice.QueryParameter(params.ReverseParam, "sort parameters, e.g. reverse=true")).
//		Param(webservice.QueryParameter(params.OrderByParam, "sort parameters, e.g. orderBy=createTime")).
//		Returns(http.StatusOK, ok, schema.PageableResponse{}))
//
//	for _, route := range webservice.Routes() {
//		logger.Debug(nil, "%s         %s", route.Method, route.Path)
//	}
//
//	c.Add(webservice)
//	return nil
//}

func addWebServiceWithStaticRoute(c *restful.Container) error {
	for _, path := range paths {
		webservice := runtime.NewWebService(path.GroupVersion)
		for _, route := range path.Routes {
			routeBuilder := webservice.Method(route.HTTPMethod).
				Path(route.Path).
				To(route.Handler).
				Doc(route.Doc).
				Metadata(route.MetaData, route.Tags).
				Writes(route.WriteDataModel)
			if route.ReadDataModel != nil {
				routeBuilder = routeBuilder.Reads(route.ReadDataModel)
			}
			routeBuilder = params.PathParameterBuilder(routeBuilder, route.PathParams)
			routeBuilder = params.QueryParameterBuilder(routeBuilder, route.QueryParams)
			for _, returnDef := range route.ReturnDefinitions {
				routeBuilder = routeBuilder.Returns(returnDef.HTTPStatus, returnDef.Message, returnDef.ReturnModel)
			}
			//Add filter in router level
			for _, filter := range route.Filter {
				routeBuilder.Filter(filter)
			}

			webservice.Route(routeBuilder)
		}
		c.Add(webservice)
	}
	return nil
}
