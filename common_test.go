package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// test helper functions will help us reduce boilerplate code when we write additional tests to test similar functionality 

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	// set Gin to test mode
	gin.SetMode(gin.TestMode)

	// run the other test functions
	os.Exit(m.Run())
}

// helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// helper function to process a request and test its response. true if the test is successful
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request
	r.ServeHTTP(w, req)
	
	if !f(w) {
		t.Fail()
	}
}

// used to store the main lists into the temporary one for testing
func saveLists() {
	tmpArticleList = articleList 
}

// used to restore the main lists from the temporary one 
func restoreLists() {
	articleList = tmpArticleList
}