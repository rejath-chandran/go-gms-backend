package main

import (
	"fmt"
	"net/http"
)

func (app *App) AllCourses(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "All courses")
}