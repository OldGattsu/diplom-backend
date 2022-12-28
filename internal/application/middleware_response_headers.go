package application

import "net/http"

func (app *Application) middlewareResponseHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("content-type", "application/json")
		// todo: add CORS headers
		handler.ServeHTTP(rw, req)
	})
}
