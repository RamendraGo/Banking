package domain

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // Use SQL Server driver

	"github.com/joho/godotenv"
)

var DB *sql.DB
var DBConnected = false

func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠ Could not load .env file, using default values")
	}

	// Fetch DB credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	log.Printf("Connecting to database %s at %s:%s with user %s\n", dbName, dbHost, dbPort, dbUser)

	// Check if any of the required environment variables are missing
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Println("⚠ Missing database configuration in environment variables")

		fmt.Println("Please set the following environment variables:")
		fmt.Println("DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME")
		DBConnected = false
		return

	}

	// Construct DSN (Data Source Name)

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, url.QueryEscape(dbPassword), dbHost, dbPort, dbName)

	// Open MySQL connection
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Println("Database connection failed:", err)
		DBConnected = false
		return
	}

	// Check if the database is reachable
	if err = DB.Ping(); err != nil {
		log.Println("Database is unreachable:", err)
		DBConnected = false
		return
	}

	DBConnected = true
	fmt.Println("Database connected successfully!")
}
