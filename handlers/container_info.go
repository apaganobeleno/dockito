package handlers

import (
	"net/http"
	"os"
)

//Home renders a Hello message
func ContainerInfo(rw http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "... IDK!"
	}
	rw.Write([]byte("Responding From: " + hostname))
}
