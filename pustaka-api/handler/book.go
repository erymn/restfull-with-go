package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": "Ery Maftiyarto",
		"desc": "Software Developer",
	})
}

func HelloHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"say":      "Hello world",
		"subtitle": "hahahahaha",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func BookHandlerV2(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":      id,
		"version": "v2",
	})
}

func Book2Handler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.IndentedJSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.IndentedJSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBookHandler(c *gin.Context) {
	//menerima 2 data, title dan price
	var newBook book.BookRequest

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
