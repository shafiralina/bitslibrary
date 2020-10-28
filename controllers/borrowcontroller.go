package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var GetBorrow = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	params := mux.Vars(r)
	id := (params["id"])

	data := models.GetBorrow(conn, id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
