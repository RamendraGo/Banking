package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RamendraGo/Banking/domain"
	"github.com/RamendraGo/Banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	// define a Gorilla multiplexer
	router := mux.NewRouter()

	// Wiring up the handler function
	// to the route

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Set up the HTTP server and route
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
