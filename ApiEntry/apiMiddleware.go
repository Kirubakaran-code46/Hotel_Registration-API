package apientry

import (
	"SDT_ADMIN_API/common"
	"SDT_ADMIN_API/helpers"
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strings"
)

type ResponseCaptureWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func (rw *ResponseCaptureWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseCaptureWriter) Write(body []byte) (int, error) {
	rw.body = append(rw.body, body...)
	return rw.ResponseWriter.Write(body)
}

func (rw *ResponseCaptureWriter) Status() int {
	if rw.status == 0 {
		return http.StatusOK
	}
	return rw.status
}

func (rw *ResponseCaptureWriter) Body() []byte {
	return rw.body
}

func APIMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lDebug := new(helpers.HelperStruct)
		lDebug.Init()
		lDebug.Log(helpers.Statement, "APIMiddleware (+)")

		(w).Header().Set("Access-Control-Allow-Origin", common.AllowedOrigin)
		(w).Header().Set("Access-Control-Allow-Credentials", "true")
		(w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "user,Content-Type, Authorization")

		if strings.EqualFold("OPTIONS", r.Method) {
			w.WriteHeader(http.StatusOK)
			return
		}

		ctx := context.WithValue(r.Context(), helpers.RequestIDKey, lDebug.Sid)
		r = r.WithContext(ctx)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		requestorDetail := GetAPIRequestDetail(lDebug, r)
		requestorDetail.Body = string(body)
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		LogRequest(lDebug, "", requestorDetail, lDebug.Sid)

		captureWriter := &ResponseCaptureWriter{ResponseWriter: w}
		next.ServeHTTP(captureWriter, r)

		lDebug.Log(helpers.Statement, "APIMiddleware (-)")
		LogResponse(lDebug, r, captureWriter.Status(), captureWriter.Body(), r.Context().Value(helpers.RequestIDKey).(string))
	})
}
