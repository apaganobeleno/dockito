package config
import (
	"dockito/handlers"
	"webo/routing"
)

//defineRoutes gets called by the generated main.go, please define routes inside it.
func DefineRoutes(r *routing.Router){
  r.Get("/", handlers.Home)
}