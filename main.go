package main

import (
	"github/HsiaoCz/leaf/etc"
	"github/HsiaoCz/leaf/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	err := etc.Init()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	router.Router(r)
	srv := http.Server{
		Handler:      r,
		Addr:         etc.Conf.App.AppPort,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	log.Fatal(srv.ListenAndServe())
}
