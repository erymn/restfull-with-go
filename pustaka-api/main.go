package main

import (
	"github.com/gin-gonic/gin"

	"pustaka-api/handler"
)

func main() {
	router := gin.Default()

	//Untuk mempermudah versioning yang pernah dibuat sebelumnya bisa pakai versioning group
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BookHandler)         //path parameter
	v1.GET("/books/:id/:title", handler.Book2Handler) //path parameter
	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBookHandler)

	v2 := router.Group("/v2")
	v2.GET("/books/:id", handler.BookHandlerV2)

	router.Run(":8880")
}
