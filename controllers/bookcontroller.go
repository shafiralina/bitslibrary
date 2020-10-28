package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	data := models.GetAllBook(conn)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetNewestBooks = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	data := models.Newest(conn)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetPopularBooks = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	data := models.Popular(conn)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	book := &models.Book{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := book.Create(conn)
	u.Respond(w, resp)
}

var GetBook = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	params := mux.Vars(r)
	id := (params["id"])

	data := models.GetBook(conn, id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	params := mux.Vars(r)
	id := (params["id"])

	data := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := data.Update(conn, id)
	u.Respond(w, resp)
}
