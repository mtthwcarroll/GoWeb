package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// structure for storing bookmarks
type bookmarksStruct struct {
	Name string
	Url  string
}

// initialize our bookmarks slice
var bookmarks []bookmarksStruct

func main() {
	// set up the server and load template(s)
	server := gin.Default()
	server.LoadHTMLGlob("templates/*.tmpl")

	// get path for the home
	server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"bookmarks": bookmarks,
		})
	})

	// post path for home
	server.POST("/", func(ctx *gin.Context) {
		bookmarks = append(bookmarks, bookmarksStruct{ctx.PostForm("name"), ctx.PostForm("url")})
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"bookmarks": bookmarks,
		})
	})
	
	// run the server
	server.Run(":8080")
}
