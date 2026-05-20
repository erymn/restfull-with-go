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

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", bookHandler)         //path parameter
	router.GET("/books/:id/:title", book2Handler) //path parameter
	router.GET("/query", queryHandler)

	router.POST("/books", postBookHandler)

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
		for _, err := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field: %s, condition: %s", err.Field(), err.Tag())
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": errorMsg,
			})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"title": newBook.Title,
		"price": newBook.Price,
	})

}
