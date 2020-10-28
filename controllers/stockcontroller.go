package controllers

import (
	"bitslibrary/models"
	u "bitslibrary/utils"
	"encoding/json"
	"net/http"
)

var CreateStock = func(w http.ResponseWriter, r *http.Request) {
	conn := models.GetDB()
	defer conn.Close()

	stock := &models.Stock{}

	err := json.NewDecoder(r.Body).Decode(stock)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := stock.Create(conn)
	u.Respond(w, resp)
}
