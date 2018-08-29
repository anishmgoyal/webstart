package errorcontroller

import (
	"net/http"

	"github.com/anishmgoyal/webstart/controllers"
)

type errorData struct {
	// ErrorCode is the http status that caused this page to render
	ErrorCode int
	// ErrorMessage is the message to display
	ErrorMessage string
}

// NotFound renders the error page with an error 404 message
func NotFound(w http.ResponseWriter, r *http.Request) {
	vd := controllers.BaseViewData(w, r)
	vd.Data = &errorData{
		ErrorCode:    404,
		ErrorMessage: "Page Not Found",
	}
	controllers.RenderView(w, "error#error", vd)
}

// InternalServerError renders the error page with an error 500 message
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	vd := controllers.BaseViewData(w, r)
	vd.Data = &errorData{
		ErrorCode:    505,
		ErrorMessage: "Internal Server Error",
	}
	controllers.RenderView(w, "error#error", vd)
}

// Forbidden renders the error page with an error 403 message
func Forbidden(w http.ResponseWriter, r *http.Request) {
	vd := controllers.BaseViewData(w, r)
	vd.Data = &errorData{
		ErrorCode:    403,
		ErrorMessage: "Forbidden",
	}
	controllers.RenderView(w, "error#error", vd)
}
