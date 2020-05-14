package schema

import "github.com/go-playground/validator/v10"

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
