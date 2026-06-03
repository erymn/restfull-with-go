package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Save(bookRequest BookRequest) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) error
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

func (s *service) Update(book Book) (Book, error) {
	return s.repository.Update(book)
}

func (s *service) Delete(book Book) error {
	return s.repository.Delete(book)
}
