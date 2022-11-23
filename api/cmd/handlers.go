package main

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"

	"github.com/igormichalak/portfolio/api/internal/models"
)

var slugRegExp = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func (app *application) blogFeedView(w http.ResponseWriter, r *http.Request) {
	posts, err := app.blogPosts.GetFeedPosts(false)
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
	slug := params.ByName("slug")
	isValidSlug := slugRegExp.MatchString(slug)
	if !isValidSlug {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	post, err := app.blogPosts.GetBySlug(slug)
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

func (app *application) snippetsFeedView(w http.ResponseWriter, r *http.Request) {
	posts, err := app.blogPosts.GetFeedPosts(true)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"snippets": posts}, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
