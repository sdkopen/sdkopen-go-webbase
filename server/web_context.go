package server

import (
	"context"
	"mime/multipart"
	"net/http"

	webBaseHttp "github.com/sdkopen/sdkopen-go-webbase/http"
)

type WebContext interface {
	Context() context.Context
	Response() http.ResponseWriter
	Request() *http.Request
	RequestHeader(key string) []string
	RequestHeaders() map[string][]string
	PathParam(key string) string
	RawQuery() string
	QueryParam(key string) string
	QueryArrayParam(key string) []string
	DecodeQueryParams(object any) error
	DecodeBody(object any) error
	DecodeFormData(object any) error
	StringBody() (string, error)
	Path() string
	Method() string
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	AddHeader(key string, value string)
	AddHeaders(headers map[string]string)
	Redirect(url string, statusCode webBaseHttp.StatusCode)
	ServeFile(path string)
	JsonResponse(statusCode webBaseHttp.StatusCode, body any)
	ErrorResponse(statusCode webBaseHttp.StatusCode, err error)
	EmptyResponse(statusCode webBaseHttp.StatusCode)
}
