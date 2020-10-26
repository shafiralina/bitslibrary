package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllUser()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := (params["id"])

	data := models.GetUser(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := (params["id"])

	data := &models.User{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := data.Update(id)
	u.Respond(w, resp)
}
