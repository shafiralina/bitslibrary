package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"net/http"
)

var Borrow = func(w http.ResponseWriter, r *http.Request) {
	borrow := &models.Borrow{}

	err := json.NewDecoder(r.Body).Decode(borrow)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := borrow.Create()
	u.Respond(w, resp)
}
