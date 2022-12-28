package application

import "net/http"

func (app *Application) middlewareIsAdmin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		u := getUserFromContext(req.Context())
		if !u.IsAdmin {
			http.Error(rw, "", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(rw, req)
	})
}
