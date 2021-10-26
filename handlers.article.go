package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
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
}