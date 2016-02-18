package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
)

type Visit struct {
	Date time.Time
}

func Count(rw http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial("mongo")
	defer session.Close()

	if err != nil {
		log.Println("Error Connecting:" + err.Error())
	}

	c := session.DB("default").C("visits")
	_ = c.Insert(Visit{Date: time.Now()})

	var count int
	count, err = c.Count()
	rw.Write([]byte("Visit Count: " + strconv.Itoa(count)))
}
