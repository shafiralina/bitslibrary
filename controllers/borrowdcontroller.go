package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"net/http"
)

var Borrowd = func(w http.ResponseWriter, r *http.Request) {

	borrowd := &models.Borrowd{}

	err := json.NewDecoder(r.Body).Decode(borrowd)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := borrowd.Create()
	u.Respond(w, resp)
}
