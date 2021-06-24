package bookshelf

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type BookShelf struct {
	Books []Book `json:"books"`
}
