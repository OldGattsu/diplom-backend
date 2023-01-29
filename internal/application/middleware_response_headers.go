package application

import "net/http"

func (app *Application) middlewareResponseHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("content-type", "application/json")
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Credentials", "true")
		rw.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		rw.Header().Add("Access-Control-Allow-Headers", "*")
		rw.Header().Add("Access-Control-Expose-Headers", "*")
		handler.ServeHTTP(rw, req)
	})
}
