package main

import (
	"github.com/gin-gonic/gin"
	"hex-go/internal/router"
)

func main() {
	r := gin.Default()
	router.App(r)
	r.Run()
}
