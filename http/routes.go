package http

import (
	"net/http"

	"github.com/oatsaysai/simple-go-http-server/app"
)

// Route is used to define http routes for the app
type Route struct {
	Name        string
	Path        string
	Method      string
	HandlerFunc func(*app.Context, http.ResponseWriter, *http.Request) error
}

// Routes is a collection of multiple http Routes
type Routes []Route
