package main

import (
	"fmt"
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

	books, err := bookRepository.FindAll()
	if err != nil {
		log.Fatal("Gagal mengambil data: ", err)
	}
	for _, b := range books {
		fmt.Printf("Book object %v", b)
	}

	bookItem, err := bookRepository.FindByID(1)
	if err != nil {
		log.Fatal("Gagal mengambil data: ", err)
	}
	fmt.Printf("Book object %v", bookItem)

	bkNew := book.Book{}
	bkNew.Title = "Belajar Go 4 (Advanced Series)"
	bkNew.Price = 160000
	bkNew.Rating = 5
	bkNew.Description = "Buku tentang belajar Go untuk advanced programmer"

	bookItem, err = bookRepository.Save(bkNew)

	if err != nil {
		log.Fatal("Gagal menyimpan data: ", err)
	}
	fmt.Printf("Book object %v", bookItem)

	//CRUD Sample
	// -----------Membuat Data------------
	// bk := book.Book{}
	// bk.Title = "Manusia Bodoh"
	// bk.Price = 120000
	// bk.Rating = 5
	// bk.Description = "Buku tentang kebodohan manusia"

	// db.Create(&bk)
	// -----------Membuat Data------------

	// // -----------Mengambil Data------------
	// var bk book.Book

	// // // Original code tanpa debug
	// // err = db.First(&bk).Error

	// //untuk debug, bisa gunakan fungsi
	// err = db.Debug().First(&bk).Error

	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// fmt.Println("Title: ", bk.Title)
	// fmt.Printf("Book object %v\n", bk)

	// //Mengambil data terakhir
	// fmt.Println("---Ambil data terakhir---")
	// err = db.Debug().Last(&bk).Error

	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// fmt.Println("Title: ", bk.Title)
	// fmt.Printf("Book object %v\n", bk)

	// //Mengambil data berdasarkan PK
	// fmt.Println("---Ambil data berdasarkan PK---")
	// err = db.Debug().First(&bk, 1).Error

	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// fmt.Println("Title: ", bk.Title)
	// fmt.Printf("Book object %v\n", bk)

	// //mengambil banyak object dan masukkan ke slice
	// fmt.Println("mengambil banyak object dan masukkan ke slice")

	// var books []book.Book
	// err = db.Debug().Find(&books).Error

	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// for _, b := range books {
	// 	fmt.Println("Title: ", b.Title)
	// 	fmt.Printf("Book object %v", b)
	// }

	// //Mengambil data menggunakan conditions
	// fmt.Println("\nMengambil data menggunakan conditions")
	// var findbooks []book.Book
	// err = db.Debug().Where("title like ?", "%Manusia%").Find(&findbooks).Error

	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// for _, b := range findbooks {
	// 	fmt.Println("Title: ", b.Title)
	// 	fmt.Printf("Book object %v", b)
	// }

	// // -----------Mengambil Data------------

	// // -----------------Update Data-------------
	// var bkUpd book.Book
	// err = db.Debug().Where("id = ?", 2).First(&bkUpd).Error
	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// bkUpd.Title = "Manusia Kera (Edisi Revisi)"
	// err = db.Save(&bkUpd).Error

	// if err != nil {
	// 	log.Fatal("Gagal update data: ", err)
	// }

	// fmt.Println("Title: ", bkUpd.Title)
	// fmt.Printf("Book object %v\n", bkUpd)

	// // -----------------Update Data-------------

	// // -----------------Delete Data-------------
	// var bkDel book.Book
	// err = db.Debug().Where("id = ?", 2).First(&bkDel).Error
	// if err != nil {
	// 	log.Fatal("Gagal mengambil data: ", err)
	// }

	// err = db.Delete(&bkDel).Error

	// //err = db.Save(&bkDel).Error

	// if err != nil {
	// 	log.Fatal("Gagal update data: ", err)
	// }

	// // -----------------Delete Data-------------

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
	//v2.GET("/findbook", handler.QueryFindBookHandler)

	router.Run(":8880")
}
