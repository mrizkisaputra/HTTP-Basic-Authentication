package utils

import (
	"net/http"
)

const USERNAME = "kiki"
const PASSWORD = "kiki"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		if !ok {
			return
		}
		if username != USERNAME && password != PASSWORD {
			return
		}
		next.ServeHTTP(writer, request)
	})
}

//type MiddlewareAuth struct {
//	Handler http.Handler
//}
//
//func (m MiddlewareAuth) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	username, password, ok := request.BasicAuth()
//	if !ok {
//		writer.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//	if username != USERNAME && password != PASSWORD {
//		writer.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//	m.Handler.ServeHTTP(writer, request)
//}
