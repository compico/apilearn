package bookshelf

import (
	"errors"
	"strconv"
)

func NewBookshelf() *BookShelf {
	bs := new(BookShelf)
	bs.Books = make([]Book, 0)
	return bs
}

func (bs *BookShelf) AddNewBook(b Book) error {
	switch {
	case b.Name == "":
		return errors.New("Field name is empty")
	case b.Author == "":
		return errors.New("Field author is empty")
	case b.Pages == 0:
		return errors.New("Field pages is empty")
	}
	b.ID = bs.GetLen()
	bs.Books = append(bs.Books, b)
	return nil
}

func (bs *BookShelf) UpdateBook(id int, b Book) error {
	err := bs.CheckBook(id)
	if err != nil {
		return err
	}
	if b.Name != "" {
		bs.Books[id].Name = b.Name
	}
	if b.Author != "" {
		bs.Books[id].Author = b.Author
	}
	if b.Pages != 0 {
		bs.Books[id].Pages = b.Pages
	}

	return nil
}

func (bs *BookShelf) GetBook(id int) (Book, error) {
	if id > (len(bs.Books) - 1) {
		return Book{}, errors.New("Book with id" + strconv.Itoa(id) + " not exits")
	}
	if id < 0 {
		return Book{}, errors.New("id less than zero")
	}
	return bs.Books[id], nil
}

func (bs *BookShelf) DeleteBook(id int) error {
	err := bs.CheckBook(id)
	if err != nil {
		return err
	}
	if id == (len(bs.Books) - 1) {
		bs.Books = bs.Books[:id-1]
		return nil
	}
	bs.Books = append(bs.Books[:id], bs.Books[id+1:]...)
	for i := id; i < len(bs.Books); i++ {
		bs.Books[i].ID -= 1
	}
	return nil
}

func (bs *BookShelf) CheckBook(id int) error {
	switch {
	case bs.Books[id] == Book{}:
		return errors.New("Book with id" + strconv.Itoa(id) + " not exits")
	case id > (len(bs.Books) - 1):
		return errors.New("Book with id" + strconv.Itoa(id) + " not exits")
	case id < 0:
		return errors.New("id less than zero")
	}
	return nil
}

func (bs *BookShelf) GetLen() int {
	return len(bs.Books)
}

func (bs *BookShelf) AddTestBooks() {
	for i := 0; i < 10; i++ {
		bs.Books = append(bs.Books, Book{
			ID:     bs.GetLen(),
			Name:   "TestName" + strconv.Itoa(i),
			Author: "TestAuthor" + strconv.Itoa(i),
			Pages:  13 + i,
		})
	}
}
