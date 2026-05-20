package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Subtitle string `json:"sub_title"`
}

func postBookHandler(c *gin.Context) {
	//menerima 2 data, title dan price
	var newBook BookInput

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"title":     newBook.Title,
		"price":     newBook.Price,
		"sub_title": newBook.Subtitle,
	})

}
