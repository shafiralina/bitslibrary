package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var CreateBorrow = func(w http.ResponseWriter, r *http.Request) {

	borrowapi := &models.BorrowAPI{}

	err := json.NewDecoder(r.Body).Decode(&borrowapi)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := borrowapi.Create()
	u.Respond(w, resp)
}

var Return = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := (params["id"])

	data := models.Return(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
