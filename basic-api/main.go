package main

import (
	"api/api/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func getEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Error while reading key %s", key)
	}

	return value
}

func getEnvVariables() map[string]string {
	env := map[string]string{
		"user":     getEnvVariable("user"),
		"password": getEnvVariable("password"),
		"dbname":   getEnvVariable("dbname"),
		"host":     getEnvVariable("host"),
		"port":     getEnvVariable("port"),
	}

	return env
}

func connectDB() {
	env := getEnvVariables()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", env["user"], env["password"], env["dbname"], env["host"], env["port"])
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

// Basic endpoint call.
func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
