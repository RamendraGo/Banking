package app

import (
	"encoding/json"
	"net/http"

	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {

	//Declare CustomerId for the request
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//pass CustomerID to the request
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}

	}

}
