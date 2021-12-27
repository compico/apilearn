package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/compico/apilearn/bookshelf"
	"github.com/julienschmidt/httprouter"
)

var (
	ct string = "Content-Type"
	aj string = "application/json"
)

func getBooksListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add(ct, aj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	j, err := json.Marshal(books.Books)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}
	_, err = w.Write(j)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}
}

func addBookHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add(ct, aj)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	p, err := strconv.Atoi(r.FormValue("pages"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", "Error convert pages count to number")
	}

	b := bookshelf.Book{
		Name:   r.FormValue("name"),
		Author: r.FormValue("author"),
		Pages:  p,
	}

	err = books.AddNewBook(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}

	w.WriteHeader(http.StatusCreated)
}

func updateBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add(ct, aj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idreq := ps.ByName("id")
	id, err := strconv.Atoi(idreq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", "Error convert pages count to number")
	}

	var p int

	if r.FormValue("pages") != "" {
		p, err = strconv.Atoi(r.FormValue("pages"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"error\":\"%v\"}", "Error convert pages count to number")
		}
	}

	b := bookshelf.Book{
		Name:   r.FormValue("name"),
		Author: r.FormValue("author"),
		Pages:  p,
	}

	err = books.UpdateBook(id, b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}
}

func getBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add(ct, aj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idreq := ps.ByName("id")
	id, err := strconv.Atoi(idreq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}

	b, err := books.GetBook(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}

	j, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}

	_, err = w.Write(j)
	if err != nil {
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", "Error convert pages count to number")
	}
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add(ct, aj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idreq := ps.ByName("id")
	id, err := strconv.Atoi(idreq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", "Error convert pages count to number")
	}

	err = books.DeleteBook(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\":\"%v\"}", err)
	}
}

func addTestBooksHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	books.AddTestBooks()
	w.WriteHeader(http.StatusOK)
}

func getLenHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add(ct, aj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	length := books.GetLen()
	fmt.Fprintf(w, "{\"_length\":%v}", strconv.Itoa(length))
}
