package interfaces

import "net/http"

type Params interface {
	ByName(name string) string
}

type HttpHandler func(http.ResponseWriter, *http.Request, Params)

type HttpHandlerWithUserID func(http.ResponseWriter, *http.Request, Params, int64)
