package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {
	router := gin.Default()

	dsn := "root:pass@word1@tcp(127.0.0.1:33060)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal connect database: ", err)
	}

	db.AutoMigrate(&book.Book{})

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
