package bootstrap

import (
	"net/http"

	"github.com/anishmgoyal/webstart/constants"
	"github.com/anishmgoyal/webstart/controllers/homecontroller"
	"github.com/anishmgoyal/webstart/resources"
)

// CreateRoutes maps URI's to the corresponding controller method
func CreateRoutes() {
	resources.MapCSSHandler()
	resources.MapImageHandler()
	resources.MapJSHandler()

	http.Handle(route("/", homecontroller.Home))
}

// Quick wrapper for StripPrefix which prevents typos
func route(path string, callback http.HandlerFunc) (string, http.Handler) {
	fn := callback
	handler := http.StripPrefix(path, fn)
	if constants.SSLEnable {
		fn = func(w http.ResponseWriter, r *http.Request) {
			if redirect, ok := constants.Domain.Map[r.Host]; ok {
				http.Redirect(w, r, redirect+r.URL.RequestURI(),
					http.StatusMovedPermanently)
				return
			}
			http.StripPrefix(path, callback).ServeHTTP(w, r)
		}
		handler = http.HandlerFunc(fn)
	}
	return path, handler
}
