package main

import (
	"github.com/compico/apilearn/bookshelf"
	"github.com/compico/apilearn/server"
)

var (
	conf  = server.NewConfig()
	books = bookshelf.NewBookshelf()
)

func init() {
	conf = &server.Server{
		Port: ":8080",
	}
	conf.InitRouter()
	conf.InitServer()
}
func main() {
	conf.R.GET("/api/v1/getbookslist", getBooksListHandler)
	conf.R.GET("/api/v1/getbook/:id", getBookHandler)
	conf.R.DELETE("/api/v1/deletebook/:id", deleteBookHandler)
	conf.R.POST("/api/v1/updatebook/:id", updateBookHandler)
	conf.R.PUT("/api/v1/addbook", addBookHandler)
	conf.R.GET("/api/v1/getlengthlist", getLenHandler)
	conf.R.PUT("/api/v1/addtestbooks", addTestBooksHandler)

	conf.S.ListenAndServe()
}
