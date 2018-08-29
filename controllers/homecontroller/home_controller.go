package homecontroller

import (
	"net/http"
	"strings"

	"github.com/anishmgoyal/webstart/controllers"
	"github.com/anishmgoyal/webstart/controllers/errorcontroller"
)

type homeData struct {
	Error      bool
	FlashTitle string
	Flash      interface{}
}

// Home handles any requests on the '/' route
func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHome(w, r)
	default:
		errorcontroller.NotFound(w, r)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	viewData := controllers.BaseViewData(w, r)

	// This is a fallback route; if there is a path in the request, it's
	// not found
	if len(r.RequestURI) > 0 && strings.Compare(r.RequestURI, "/") != 0 {
		errorcontroller.NotFound(w, r)
		return
	}

	controllers.RenderView(w, "home#index", viewData)
}
