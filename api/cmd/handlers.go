package main

import "net/http"

func blogFeed(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog feed"))
}

func blogPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("single blog post"))
}
