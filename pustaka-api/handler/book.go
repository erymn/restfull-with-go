package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// // Untuk memindahkan service dari main.go yang perlu dilakukan:
// 1. Buat struct untuk book handler
type bookHandler struct {
	bookService book.Service
}

//  2. Buat function NewBookHandler yang akan memiliki parameter book.Service
//     dengan return value sebagai bookHandler struct
func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// // function RootHandler, supaya bisa dimiliki oleh bookHandler, maka function dibuat method
// // dengan menambahkan nama handler didepan function name nya
// func (h *bookHandler) RootHandler(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"name": "Ery Maftiyarto",
// 		"desc": "Software Developer",
// 	})
// }

// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"say":      "Hello world",
// 		"subtitle": "hahahahaha",
// 	})
// }

// func (h *bookHandler) BookHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"id": id,
// 	})
// }

// func (h *bookHandler) BookHandlerV2(c *gin.Context) {
// 	id := c.Param("id")
// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"id":      id,
// 		"version": "v2",
// 	})
// }

// func (h *bookHandler) Book2Handler(c *gin.Context) {
// 	id := c.Param("id")
// 	title := c.Param("title")
// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})
// }

// func (h *bookHandler) QueryHandler(c *gin.Context) {
// 	title := c.Query("title")
// 	price := c.Query("price")

// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"title": title,
// 		"price": price,
// 	})
// }

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookList := []book.BookResponse{}
	for _, b := range books {
		bookList = append(bookList, book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Price:       b.Price,
			Rating:      b.Rating,
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": bookList,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
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

	bookInp, err := h.bookService.Save(newBook)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": bookInp,
	})

}
