package main

func initializeRoutes() {
	// handle the index route
	router.GET("/", showIndexPage)

	// handle GET requests at /article/view/some_article_id
	// stores the value of the last part of the route in the route paramenter named article_id which we can access in the route handler
	router.GET("/article/view/:article_id", getArticle)
}