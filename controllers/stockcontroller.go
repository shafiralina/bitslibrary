package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"net/http"
)

var CreateStock = func(w http.ResponseWriter, r *http.Request) {

	stock := &models.Stock{}

	err := json.NewDecoder(r.Body).Decode(stock)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := stock.Create()
	u.Respond(w, resp)
}
