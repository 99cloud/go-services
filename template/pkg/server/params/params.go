package params

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/schema"

	"github.com/emicklei/go-restful"
)

const (
	PagingParam     = "paging"
	OrderByParam    = "orderBy"
	ConditionsParam = "conditions"
	ReverseParam    = "reverse"
)

type Conditions struct {
	Match map[string]string
	Fuzzy map[string]string
}

func ParsePaging(paging string) (limit, offset int) {

	limit = 10
	offset = 0
	if groups := regexp.MustCompile(`^limit=(-?\d+),page=(\d+)$`).FindStringSubmatch(paging); len(groups) == 3 {
		limit, _ = strconv.Atoi(groups[1])
		page, _ := strconv.Atoi(groups[2])
		offset = (page - 1) * limit
	}
	return
}

func ParseConditions(conditionsStr string) (*Conditions, error) {

	conditions := &Conditions{Match: make(map[string]string, 0), Fuzzy: make(map[string]string, 0)}

	if conditionsStr == "" {
		return conditions, nil
	}

	// ?conditions=key1=value1,key2~value2,key3=
	for _, item := range strings.Split(conditionsStr, ",") {
		// exact query: key=value, if value is empty means label value must be ""
		// fuzzy query: key~value, if value is empty means label value is "" or label key not exist
		if groups := regexp.MustCompile(`(\S+)([=~])(\S+)?`).FindStringSubmatch(item); len(groups) >= 3 {
			value := ""

			if len(groups) > 3 {
				value = groups[3]
			}

			if groups[2] == "=" {
				conditions.Match[groups[1]] = value
			} else {
				conditions.Fuzzy[groups[1]] = value
			}
		} else {
			return nil, fmt.Errorf("invalid conditions")
		}
	}
	return conditions, nil
}

func ParseReverse(req *restful.Request) bool {
	reverse := req.QueryParameter(ReverseParam)
	b, err := strconv.ParseBool(reverse)
	if err != nil {
		return false
	}
	return b
}

func GetStringValueWithDefault(req *restful.Request, name string, dv string) string {
	v := req.QueryParameter(name)
	if v == "" {
		v = dv
	}
	return v
}

func PathParameterBuilder(rb *restful.RouteBuilder, paras []schema.Parameter) *restful.RouteBuilder {
	for _, para := range paras {
		rb = rb.Param(restful.PathParameter(para.Name, para.Description).Required(para.Required).DefaultValue(para.DefaultValue))
	}
	return rb
}

func QueryParameterBuilder(rb *restful.RouteBuilder, paras []schema.Parameter) *restful.RouteBuilder {
	for _, para := range paras {
		rp := restful.QueryParameter(para.Name, para.Description).Required(para.Required).DefaultValue(para.DefaultValue)
		if para.AllowMultiple == true {
			rp.AllowMultiple(true)
		}
		if para.DataType != "" {
			rp.DataType(para.DataType)
		}
		rb = rb.Param(rp)
	}
	return rb
}
