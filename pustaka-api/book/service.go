package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Save(bookRequest BookRequest) (Book, error)
	Update(id int, bookRequest BookRequest) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	return book, err
}

func (s *service) Save(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()

	// ubah bookrequest ke book
	newBook := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}

	bookInput, err := s.repository.Save(newBook)
	return bookInput, err
	//return s.repository.Save(newBook)
}

func (s *service) Update(id int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByID(id)
	if err != nil {
		return book, err
	}

	price, _ := bookRequest.Price.Int64()

	// ubah bookrequest ke book
	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = bookRequest.Rating

	bookInput, err := s.repository.Update(book)
	return bookInput, err
	//return s.repository.Save(newBook)
}

func (s *service) Delete(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	if err != nil {
		return book, err
	}
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
