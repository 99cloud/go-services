package database

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/constants"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/httputils"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/logger"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/schema"
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/server/errors"
)

// recover from panic and log runtime stack trace
func CatchError() {
	if err := recover(); err != nil {
		var stackMsg [4096]byte
		runtime.Stack(stackMsg[:], false)
		logger.Debug(nil, string(stackMsg[:]))
	}
}

// HttpDataHandle is a interface encapsulates common processing methods of HTTP requests.
// Http Request can be completed by a HttpDataHandle object.
type HttpDataHandle struct {
	req       *restful.Request
	resp      *restful.Response
	db        *gorm.DB // db is a pointer of initialed gorm.DB, which is used to perform Find, Save, .etc.
	tmpDb     *gorm.DB // temDb is the output of Find, Save, .etc.
	uri       string
	errorCode int    // http response code
	errorMsg  string // http response msg
	// errorInfoLevel is used in function SetErrorCodeMsg,
	// when input level is larger than this, errorCode and errorMsg will be override.
	errorInfoLevel int8
	PostbackUri    string // postback interface
	returned       bool   // indicates whether response has been returned. response can only be returned once.
}

// generate a object of HttpDataHandle
func NewHttpDataHandle(req *restful.Request, resp *restful.Response) *HttpDataHandle {
	h := new(HttpDataHandle)
	if req == nil || resp == nil {
		panic(errors.New("req or resp is nil"))
	} else {
		h.req = req
		h.resp = resp
	}
	// Because resp is not nil, error response can be returned.
	defer h.handlePanic()
	h.setDefaultErrorCodeMsg()
	h.uri = req.Request.RequestURI
	db, err := GetGorm()
	if err != nil {
		logger.Error(nil, err.Error())
		panic(err)
	}
	h.db = db
	h.tmpDb = h.db.New()
	return h
}

// set default errorCode, errorMsg and errorInfoLevel.
func (h *HttpDataHandle) setDefaultErrorCodeMsg() {
	h.errorCode = http.StatusInternalServerError
	h.errorMsg = constants.ERR_MSG_INTERNAL_ERR
	h.errorInfoLevel = 0
}

// used to custom errorCode and errorMsg in response
// if level is larger than h.errorInfoLevel, h.errorCode and h.errorMsg will be override.
func (h *HttpDataHandle) SetErrorCodeMsg(code int, msg string, level int8) {
	if level >= h.errorInfoLevel {
		h.errorCode = code
		h.errorMsg = msg
		h.errorInfoLevel = level
	}
}

func (h *HttpDataHandle) handlePanic() {
	if err := recover(); err != nil {
		if !h.returned {
			h.returned = true
			httputils.WriteCommonResponse(h.resp, h.errorCode, h.errorMsg, h.uri)
		}
		h.setDefaultErrorCodeMsg()
		panic(err)
	}
	h.setDefaultErrorCodeMsg()
}

// collect errors after performing Find, Save, .etc.
func (h *HttpDataHandle) dbPanic() {
	if len(h.tmpDb.GetErrors()) > 0 {
		err := h.tmpDb.GetErrors()[0]
		logger.Error(nil, err.Error())
		if h.tmpDb.RecordNotFound() {
			h.SetErrorCodeMsg(http.StatusNotFound, constants.ERR_MSG_REQID_NOT_FOUND, 0)
		}
		h.tmpDb = h.db.New()
		panic(err)
	}
	h.tmpDb = h.db.New()
	h.setDefaultErrorCodeMsg()
}

// validate a object according to Struct Validator Tag and Custom Validator Functions.
// Custom Validator Functions are specified by customFieldValidators and customStructValidators.
func (h *HttpDataHandle) Validate(s interface{}, customFieldValidators map[string]validator.Func,
	customStructValidators []schema.StructLevelValidation) {
	defer h.handlePanic()
	validate := validator.New()
	for k, v := range customFieldValidators {
		_ = validate.RegisterValidation(k, v)
	}
	for _, v := range customStructValidators {
		validate.RegisterStructValidation(v.StructValidation, v.StructType)
	}
	err := validate.Struct(s)
	if err != nil {
		logger.Error(nil, err.Error())
		h.SetErrorCodeMsg(http.StatusBadRequest, constants.ERR_REQUIRED_ITEMS_INVALID, 0)
		panic(err)
	}
}

// parse the JSON-encoded data and store the result
func (h *HttpDataHandle) Unmarshal(bytes []byte, output interface{}) {
	defer h.handlePanic()
	err := json.Unmarshal(bytes, output)
	if err != nil {
		logger.Error(nil, err.Error())
		// default errorCode and errorMsg.
		// there is no need to invoke SetErrorCodeMsg
		panic(err)
	}
}

// return the JSON encoding of v.
func (h *HttpDataHandle) Marshal(v interface{}, output *[]byte) {
	defer h.handlePanic()
	bytes, err := json.Marshal(v)
	if err != nil {
		logger.Error(nil, err.Error())
		// default errorCode and errorMsg.
		panic(err)
	}
	// copy bytes to output
	*output = bytes
}

// execute a custom function and handle various errors, which can decrease repeat codes.
func (h *HttpDataHandle) Execute(errorCode int, errorMsg string, f func(...interface{}) error, params ...interface{}) {
	defer h.handlePanic()
	err := f(params...)
	if err != nil {
		logger.Error(nil, err.Error())
		h.SetErrorCodeMsg(errorCode, errorMsg, math.MaxInt8)
		panic(err)
	}
}

// panic a costumed error
func (h *HttpDataHandle) RaiseError(errorCode int, errorMsg string, err error) {
	defer h.handlePanic()
	if err == nil {
		err = errors.New(errorMsg)
	}
	logger.Error(nil, err.Error())
	h.SetErrorCodeMsg(errorCode, errorMsg, math.MaxInt8)
	panic(err)
}

// check the Accept header and read the content into the entityPointer.
func (h *HttpDataHandle) ReadEntity(entityPointer interface{}) {
	defer h.handlePanic()
	err := h.req.ReadEntity(entityPointer)
	if err != nil {
		// log error, return error response
		logger.Error(nil, err.Error())
		h.SetErrorCodeMsg(http.StatusBadRequest, constants.ERR_MSG_INTERNAL_ERR, 0)
		panic(err)
	}
}

// transform a object from one struct to another.
func (h *HttpDataHandle) Transformation(f func(interface{}, interface{}) error, input interface{}, output interface{}) {
	defer h.handlePanic()
	err := f(input, output)
	if err != nil {
		logger.Error(nil, err.Error())
		// log error, return error response
		panic(err)
	}
}

// response a http request normally
func (h *HttpDataHandle) NormalReturn(output interface{}) {
	if output == nil {
		h.resp.WriteHeader(http.StatusNoContent)
	} else {
		err := h.resp.WriteHeaderAndJson(http.StatusOK, output, "application/json")
		if err != nil {
			logger.Error(nil, err.Error())
			panic(err)
		}
	}
}

// set search conditions.
// support performing 'func Where' multi times, e.g.,
// h.Where(.., ..)
// h.Where(.., ..)
// ...
// h.Find(..)
func (h *HttpDataHandle) Where(query interface{}, args ...interface{}) {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Where(query, args)
	// There is no need to invoke Function dbPanic,
	// because dbPanic will reset tmpDb, make Function Where no valuable.
}

func (h *HttpDataHandle) Create(value interface{}) {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Create(value)
	h.dbPanic()
}

func (h *HttpDataHandle) Delete(value interface{}, where ...interface{}) {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Delete(value, where...)
	h.dbPanic()
}

func (h *HttpDataHandle) ForceDelete(value interface{}, where ...interface{}) {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Unscoped().Delete(value, where...)
	h.dbPanic()
}

// find records from database.
// if pagination=true, page_size and page_number in h.req, pagination will be performed.
func (h *HttpDataHandle) Find(pagination bool, out interface{}, where ...interface{}) {
	defer h.handlePanic()
	if pagination {
		parasString := []string{"page_size", "page_number"}
		var parasInt []int
		errorMsg := "pagesize or page is not correct."
		pattern := "\\d+"
		for _, para := range parasString {
			value := h.req.QueryParameter(para)
			if value == "" {
				break
			}
			result, err := regexp.MatchString(pattern, value)
			if !result || err != nil {
				logger.Error(nil, errorMsg)
				panic(errors.New(errorMsg))
			}
			paraInt, err := strconv.Atoi(value)
			if paraInt == 0 || err != nil {
				logger.Error(nil, errorMsg)
				panic(errors.New(errorMsg))
			}
			parasInt = append(parasInt, paraInt)
		}
		if len(parasInt) == 2 {
			var count int
			if reflect.TypeOf(out).Kind().String() != "ptr" ||
				reflect.TypeOf(out).Elem().Kind().String() != "slice" {
				typeErrMsg := "argument type error"
				logger.Error(nil, typeErrMsg)
				panic(errors.New(typeErrMsg))
			}
			tmpValue := reflect.New(reflect.TypeOf(out).Elem())
			h.tmpDb.Find(tmpValue.Interface(), where...).Count(&count)
			h.resp.AddHeader("counts", strconv.Itoa(count))
			h.tmpDb = h.tmpDb.Limit(parasInt[0]).Offset((parasInt[1] - 1) * parasInt[0])
		}
	}
	h.tmpDb = h.tmpDb.Find(out, where...)
	h.dbPanic()
}

func (h *HttpDataHandle) First(out interface{}, where ...interface{}) {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.First(out, where...)
	h.dbPanic()
}

func (h *HttpDataHandle) Update(oldValue interface{}, newValue interface{}) int64 {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Model(oldValue).Update(newValue)
	rowsAffected := h.tmpDb.RowsAffected
	h.dbPanic()
	return rowsAffected
}

func (h *HttpDataHandle) UpdateColumns(oldValue interface{}, newValue interface{}) int64 {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Model(oldValue).UpdateColumns(newValue)
	rowsAffected := h.tmpDb.RowsAffected
	h.dbPanic()
	return rowsAffected
}

func (h *HttpDataHandle) Save(value interface{}) {
	defer h.handlePanic()
	h.tmpDb = h.tmpDb.Save(value)
	h.dbPanic()
}

func (h *HttpDataHandle) Transaction(f func(db *gorm.DB) error) {
	defer h.handlePanic()
	err := h.tmpDb.Transaction(f)
	if err != nil {
		panic(err)
	}
	h.dbPanic()
}
