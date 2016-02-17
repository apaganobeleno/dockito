package config

import (
	"dockito/handlers"
	"webo/routing"
)

//defineRoutes gets called by the generated main.go, please define routes inside it.
func DefineRoutes(r *routing.Router) {
	r.Get("/", handlers.Hello)
	r.Get("/count", handlers.Count)
	r.Get("/hello", handlers.Hello)
	r.Get("/who_are_you", handlers.ContainerInfo)
}
