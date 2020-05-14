package server

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/pkg/logger"
	"bytes"
	"fmt"
	"net/http"
	"runtime"
)

func LogStackOnRecover(panicReason interface{}, httpWriter http.ResponseWriter) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("recover from panic situation: - %v\r\n", panicReason))
	for i := 2; ; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
	}
	logger.Error(nil, buffer.String())
	httpWriter.WriteHeader(http.StatusInternalServerError)
	httpWriter.Write([]byte("recover from panic situation"))
}
