package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	//Untuk mempermudah versioning yang pernah dibuat sebelumnya bisa pakai versioning group
	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id", bookHandler)         //path parameter
	v1.GET("/books/:id/:title", book2Handler) //path parameter
	v1.GET("/query", queryHandler)

	v1.POST("/books", postBookHandler)

	v2 := router.Group("/v2")
	v2.GET("/books/:id", bookHandlerV2)

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

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func bookHandlerV2(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":      id,
		"version": "v2",
	})
}

func book2Handler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.IndentedJSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {
	//menerima 2 data, title dan price
	var newBook BookInput

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		errorMsgs := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			//Capture all error yang dilooping
			errorMsg := fmt.Sprintf("Error on field: %s, condition: %s", err.Field(), err.Tag())
			errorMsgs = append(errorMsgs, errorMsg)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMsgs,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"title": newBook.Title,
		"price": newBook.Price,
	})

}
