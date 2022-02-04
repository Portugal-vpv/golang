package middleware

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"

	"api/api/models"
	"fmt"
	"log"
	_ "log"
	"net/http"
	_ "os"
	_ "strconv"

	_ "github.com/gorila/mux"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type response struct {
	ID      int64  `json:"id.omitempty"`
	Message string `json;"message.omitempty"`
}

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

func connectDB() *sql.DB {
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
	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)

	}

	insertID := insertUser(user)

	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func insertUser(user models.User) int64 {
	db := connectDB()

	defer db.Close()

	sqlStatement := `INSERT INTO users(name, location, age) VALUES ($1, $2, $2) RETURNING userid`
	var id int64

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}
