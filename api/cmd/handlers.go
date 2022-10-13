package main

import (
	"net/http"
)

func (app *application) blogFeedView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog feed"))
}

func (app *application) blogPostView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("single blog post"))
}

func (app *application) blogTagsView(w http.ResponseWriter, r *http.Request) {
	tags, err := app.blogTags.GetAll()
	if err != nil {
		app.serverError(w, err)
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"tags": tags}, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
