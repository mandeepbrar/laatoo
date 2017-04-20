package net

import (
	"net/http"
)

type HandlerFunc func(WebContext) error

type Router interface {
	//Get a sub router
	Group(path string) Router
	//Use middleware
	Use(handler HandlerFunc)
	Get(path string, handler HandlerFunc)
	Options(path string, handler HandlerFunc)
	Post(path string, handler HandlerFunc)
	Put(path string, handler HandlerFunc)
	Delete(path string, handler HandlerFunc)
	UseMW(handler func(http.Handler) http.Handler)
	UseMiddleware(handler http.HandlerFunc)
}
