package controllers

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ViewData defines the base properties used by a view
type ViewData struct {
	Data       interface{}
	Constants  map[string]interface{}
	CurrentURI string
	header     http.Header
}

// Base contains all variables shared by controllers
var Base struct {
	// Templates contains all views available to controllers
	Templates map[string]*template.Template
	// DB is the database connection to be used by controllers and passed to models
	Db *sql.DB
}

// BaseInitialization initializes all controllers
func BaseInitialization(templates map[string]*template.Template, db *sql.DB) {
	Base.Templates = templates
	Base.Db = db
}

// BaseViewData gets any fields necessary for rendering a basic view
func BaseViewData(w http.ResponseWriter, r *http.Request) ViewData {
	return ViewData{
		Data:       nil,
		Constants:  map[string]interface{}{},
		CurrentURI: r.RequestURI,
		header:     r.Header,
	}
}

// RenderView attempts to render a view. Gives a 500 error on failure
func RenderView(w http.ResponseWriter, templateName string, data ViewData) {
	var buff bytes.Buffer

	tmpl, ok := Base.Templates[templateName]
	if !ok {
		log.Printf("Unknown Template: %s", templateName)
		http.Error(w, "Failed to render view", http.StatusInternalServerError)
		return
	}

	err := tmpl.ExecuteTemplate(&buff, "base", data)
	if err != nil {
		log.Printf("Failed to render template: %s", err.Error())
		http.Error(w, "Failed to render view", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if strings.Contains(data.header.Get("Accept-Encoding"), "gzip") {
		gz := gzip.NewWriter(w)
		defer gz.Close()
		w.Header().Set("Content-Encoding", "gzip")
		gz.Write(buff.Bytes())
	} else {
		w.Write(buff.Bytes())
	}
}

// RenderJSON attempts to render an object as JSON. Gives a 500 error on failure
func RenderJSON(w http.ResponseWriter, value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		log.Printf("Failed to render json: %s", err.Error())
		http.Error(w, "{\"error\": \"failed\"}", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// RenderJSONWithStatus attempts to render an object as JSON with the
// specified status code, failing with 500 if not possible
func RenderJSONWithStatus(w http.ResponseWriter, value interface{}, status int) {
	b, err := json.Marshal(value)
	if err != nil {
		log.Printf("Failed to render json: %s", err.Error())
		http.Error(w, "{\"error\": \"failed\"}", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	http.Error(w, string(b), status)
}
