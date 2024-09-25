package models

import (
	"net/http"
)

// model for creating a new route
type Route struct {
	Path   string
	View   string
	Method string
}

// model for making a new fetch handler
type FetchData struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
	Method  string
}

// model used in the sql package
type Value struct {
	Data string
}
