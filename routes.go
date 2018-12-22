// routes.go

package main

func initializeRoutes() {
	// handle the index route - which will display the list of articles
	router.GET("/", showIndexPage)

	// handle GET requests to individual articles
	router.GET("/a/view/:article_id", getArticle)
}
