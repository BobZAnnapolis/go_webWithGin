// main.go

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// set the router as the default one provided by Gin
	router = gin.Default()

	// process the templates at the start so that they don't have to be loaded
	// from the disk again, ie, speed up HTML page rendering
	router.LoadHTMLGlob("templates/*")

	// initialize the routes
	initalizeRoutes()

	// start serving the application
	router.Run()
}
