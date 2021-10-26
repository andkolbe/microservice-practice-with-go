package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// call render function with the same name of the template to render
	render(c, gin.H{"title": "Home Page", "payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the template
			render(c, gin.H{"title": article.Title, "payload": article}, "article.html")
		} else {
			// if the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
