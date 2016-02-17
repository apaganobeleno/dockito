package main
import (
  "log"
  webo "webo/server"
  "dockito/config"
)

func main(){
	s := webo.NewServer(config.DefineRoutes)
  s.Start("8080")
	defer log.Println("| dockito Closing")
}
