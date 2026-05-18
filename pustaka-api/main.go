package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)

	router.Run(":8880")
}

func rootHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": "Ery Maftiyarto",
		"desc": "Software Developer",
	})
}

func helloHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"say":      "Hello world",
		"subtitle": "hahahahaha",
	})
}
