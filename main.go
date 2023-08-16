package main

import (
	"github/HsiaoCz/leaf/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	addr = "127.0.0.1:9091"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.Router(r)
	srv := http.Server{
		Handler:      r,
		Addr:         addr,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
