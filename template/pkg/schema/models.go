package schema

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
