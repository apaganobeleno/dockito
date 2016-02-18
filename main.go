package main

import (
	"dockito/config"
	"log"
	webo "webo/server"
)

func main() {
	s := webo.NewServer(config.DefineRoutes)
	s.Start("8088")
	defer log.Println("| Dockito Closing")
}
