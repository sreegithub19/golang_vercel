package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

// CREATE ENDPOINT

func myRoute(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, `
		Hello admin from golang in vercel
		`)
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
		Hello from golang in vercel
		`)
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)

}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
