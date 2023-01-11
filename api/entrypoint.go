package handler

// package main -> for running locally

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func Handler() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
        <html><h1>Hi ping</h1></html>
        `))
	})
	r.GET("/pong", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
        <html><h1>Hi pong</h1></html>
        `))
	})
	r.Run() // listen and serve on http://localhost:8080/ping
}
