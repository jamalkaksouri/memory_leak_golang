package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	pprof.Register(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("failed to running app")
	}
}
