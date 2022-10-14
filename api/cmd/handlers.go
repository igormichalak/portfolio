package main

import (
	"net/http"
)

func (app *application) blogFeedView(w http.ResponseWriter, r *http.Request) {
	posts, err := app.blogPosts.GetFeedPosts()
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"posts": posts}, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) blogPostView(w http.ResponseWriter, r *http.Request) {
	post, err := app.blogPosts.Get(0)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"post": post}, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) blogTagsView(w http.ResponseWriter, r *http.Request) {
	tags, err := app.blogTags.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"tags": tags}, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
