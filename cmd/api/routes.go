package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *App) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/course", app.AllCourses)

	return router
}
