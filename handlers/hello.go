package handlers

import "net/http"

//Home renders a Hello message
func Hello(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello From Webo!"))
}
