package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	router.GET("/", func(c *gin.Context) { // Context contains all the information about the request that the handler might need to process it
		// call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (ok)
			http.StatusOK,
			// use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H {
				"title": "Home Page", // set the title to Home Page 
			},
		)
	})

	// start serving the application
	router.Run()
}