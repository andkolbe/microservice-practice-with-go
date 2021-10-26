package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// initialize the routes
	initializeRoutes()

	// start serving the application
	router.Run()
}

// renders HTML, JSON, or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
	

}