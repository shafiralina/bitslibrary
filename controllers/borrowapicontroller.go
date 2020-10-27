package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
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
