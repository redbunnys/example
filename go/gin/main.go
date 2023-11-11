package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-gin/route"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func main() {
	flag.Parse()

	r := gin.New()
	route.RouteInit(r)

	server := http.Server{
		Addr: *addr,
	}
	server.Handler = r
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
