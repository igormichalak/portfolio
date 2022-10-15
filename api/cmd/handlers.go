package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/igormichalak/portfolio/api/internal/models"
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
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	post, err := app.blogPosts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
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
