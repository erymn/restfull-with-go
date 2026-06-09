package book

import "gorm.io/gorm"

// Cara implementasi Repository class di golang
// 1. Implementasi Interface untuk repository
type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Save(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

// 2. Membuat struct agar interface bisa digunakan dalam bentuk struct (representasi dari class)
type repository struct {
	// karena butuh DB gorm
	db *gorm.DB
}

// 3. Membuat instance dari struct repository
// -> diibaratkan sebagai class constuctor yang akan memiliki return struct repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// 4. Implementasi method dari interface
func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) FindByID(id int) (Book, error) {
	var book Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Save(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error
	if err != nil {
		return Book{}, err
	}

	return book, nil
}
