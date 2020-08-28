package schema

import "github.com/go-playground/validator/v10"

type PageableResponse struct {
	Items      interface{} `json:"items" description:"paging data"`
	TotalCount int         `json:"total_count" description:"total count"`
}

type Parameter struct {
	Required      bool
	DataFormat    string
	DefaultValue  string
	Name          string
	Description   string
	DataType      string
	AllowMultiple bool
}

type CommonResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

type StructLevelValidation struct {
	StructType       interface{}
	StructValidation validator.StructLevelFunc
}
