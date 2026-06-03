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

	bookRepository := book.NewRepository(db)

	// Mekanisme lewat Service baru kemudian lewat repository
	bookService := book.NewService(bookRepository)

	// buat bookHandler/controller
	bookHandler := handler.NewBookHandler(bookService)

	//Untuk mempermudah versioning yang pernah dibuat sebelumnya bisa pakai versioning group
	v1 := router.Group("/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id", bookHandler.BookHandler)         //path parameter
	// v1.GET("/books/:id/:title", bookHandler.Book2Handler) //path parameter
	// v1.GET("/query", bookHandler.QueryHandler)

	v1.POST("/books", bookHandler.PostBookHandler)
	v1.GET("/books", bookHandler.GetBooksHandler)

	//v2 := router.Group("/v2")
	//v2.GET("/books/:id", bookHandler.BookHandlerV2)
	//v2.GET("/findbook", handler.QueryFindBookHandler)

	router.Run(":8880")
}

//Layer Aplikasi
// 1. main
// 2. handler
// 3. Services
// 4. Repository
// 5. DB (GORM)
// 6. MySQL
