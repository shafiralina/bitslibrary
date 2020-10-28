package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllBook()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetNewestBooks = func(w http.ResponseWriter, r *http.Request) {
	data := models.Newest()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetPopularBooks = func(w http.ResponseWriter, r *http.Request) {
	data := models.Popular()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateBook = func(w http.ResponseWriter, r *http.Request) {

	book := &models.Book{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := book.Create()
	u.Respond(w, resp)
}

var GetBook = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := (params["id"])

	data := models.GetBook(id)
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
