package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func gotDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv(key)
}

func set envVariables(key string) bool {

}


func main() {
	connStr := "user=postgres password=postgres dbname=postgres host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected to db!")
}
