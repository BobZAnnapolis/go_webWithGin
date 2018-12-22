// routes.go

package main

func initializeRoutes() {
	// handle the index route - which will display the list of articles
	router.GET("/", showIndexPage)
}
