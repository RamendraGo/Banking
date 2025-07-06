package app

import (
	"encoding/json"
	"net/http"

	"github.com/RamendraGo/Banking/dto"
	"github.com/RamendraGo/Banking/service"
	"github.com/gorilla/mux"
)

type TransactionHandlers struct {
	service service.TransactionService
}

func (h TransactionHandlers) NewTransaction(w http.ResponseWriter, r *http.Request) {

	//Declare AccountId for the request
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.NewTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//pass CustomerID to the request
		request.AccountId = accountId
		request.CustomerId = customerId

		transaction, appError := h.service.NewTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, transaction)
		}

	}

}
