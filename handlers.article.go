package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (ok)
		http.StatusOK,
		// use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H {
			"title": "Home Page", // set the title to Home Page 
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {
	// check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// call the HTML method of the Context to render a template
			c.HTML(
				// set the HTTP status code to 200 (OK)
				http.StatusOK,
				// use the index.html template
				"article.html",
				// pass the data that the page uses
				gin.H{
					"title": article.Title,
					"payload": article,
				},
			)
		} else {
			// if the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}