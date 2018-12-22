// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// set the router as the default one provided by Gin
	router = gin.Default()

	// process the templates at the start so that they don't have to be loaded
	// from the disk again, ie, speed up HTML page rendering
	router.LoadHTMLGlob("templates/*")

	// define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. later on we'll create
	// standalone functions that will be used a route handlers
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// set the HTTP status to 200
			http.StatusOK,
			// use the index.html template
			"index.html",
			// pass the data that the page uses, ie, title
			gin.H{
				"title": "Window Title String",
			},
		)
	})

	// start serving the application
	router.Run()
}
