package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go",
			"tar":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run()
}
