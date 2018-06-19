package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetPost",
		"GET",
		"blog/{id}",
		GetPost,
	},
	Route{
		"CreatePost",
		"POST",
		"blog/new",
		CreatePost,
	},
	Route{
		"DeletePost",
		"DELETE",
		"/blog/{id}",
		DeletePost,
	},
	Route{
		"EditPost",
		"PUT",
		"/blog/{id}",
		EditPost,
	},
	Route{
		"GetAllPosts",
		"GET",
		"/blog",
		GetAllPosts,
	},
}
