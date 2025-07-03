package domain

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/RamendraGo/Banking/logger"
	_ "github.com/denisenkom/go-mssqldb" // Use SQL Server driver
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/joho/godotenv"
)

var DB *sqlx.DB
var DBConnected = false

func checkMissingEnvVars(dbUser, dbPassword, dbHost, dbPort, dbName string) []string {
	var missing []string
	if dbUser == "" {
		missing = append(missing, "DB_USER")
	}
	if dbPassword == "" {
		missing = append(missing, "DB_PASSWORD")
	}
	if dbHost == "" {
		missing = append(missing, "DB_HOST")
	}
	if dbPort == "" {
		missing = append(missing, "DB_PORT")
	}
	if dbName == "" {
		missing = append(missing, "DB_NAME")
	}
	return missing
}

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

	// Connecting to database is now logged using logger.Info above.

	logger.Info("Connecting to database",
		zap.String("dbName", dbName),
		zap.String("dbHost", dbHost),
		zap.String("dbPort", dbPort),
		zap.String("dbUser", dbUser),
	)

	// Check if any of the required environment variables are missing
	if missingEnvVars := checkMissingEnvVars(dbUser, dbPassword, dbHost, dbPort, dbName); len(missingEnvVars) > 0 {
		log.Println("⚠ Missing database configuration in environment variables")

		fmt.Println("Please set the following environment variables:")
		for _, v := range missingEnvVars {
			fmt.Println(v)
		}
		DBConnected = false
		return

	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, url.QueryEscape(dbPassword), dbHost, dbPort, dbName)

	// Open MySQL connection
	DB, err = sqlx.Open("sqlserver", dsn)
	if err != nil {
		logger.Info("Database connection failed", zap.Error(err))
		DBConnected = false
		return
	}

	// Check if the database is reachable
	if err = DB.Ping(); err != nil {

		logger.Info("Database is unreachable")

		DBConnected = false
		return
	}

	DBConnected = true
	logger.Info("Database connected successfully!")

}

func GetDB() *sqlx.DB {
	return DB
}
