package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RamendraGo/Banking/domain"
	"github.com/RamendraGo/Banking/logger"
	"github.com/RamendraGo/Banking/service"
	"github.com/gorilla/mux"
)

func checkMissingEnvVars(address, port string) []string {
	var missing []string
	if address == "" {
		missing = append(missing, "SERVER_ADDRESS")
	}
	if port == "" {
		missing = append(missing, "SERVER_PORT")
	}

	return missing
}

func Start() {

	// define a Gorilla multiplexer
	router := mux.NewRouter()
	domain.Connect() // initializes domain.DB

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// Check if any of the required environment variables are missing
	if missingEnvVars := checkMissingEnvVars(address, port); len(missingEnvVars) > 0 {
		logger.Error("⚠ Missing configuration in environment variables")

		logger.Info("Please set the following environment variables:")
		for _, v := range missingEnvVars {
			fmt.Println(v)
		}
		return
	}

	// Wiring up the handler function
	// to the route

	// Use the stub repository for testing
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Use the database repository for production
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(domain.DB))}

	// Set up the HTTP server and route
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	fmt.Printf("Starting server on %s:%s\n", address, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
