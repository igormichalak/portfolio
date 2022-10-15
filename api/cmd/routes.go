package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	router.HandlerFunc(http.MethodGet, "/v1/blog/feed", app.blogFeedView)
	router.HandlerFunc(http.MethodGet, "/v1/blog/post/:id", app.blogPostView)
	router.HandlerFunc(http.MethodGet, "/v1/blog/tags", app.blogTagsView)

	standard := alice.New(app.recoverPanic, app.logRequest)

	return standard.Then(router)
}
