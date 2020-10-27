package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var GetDetailBorrow = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := (params["id"])

	data := models.GetDetailBorrow(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
