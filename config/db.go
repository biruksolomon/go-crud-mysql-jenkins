package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, using system env variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database not reachable:", err)
	}

	log.Println("✅ MySQL Connected Successfully")
}
