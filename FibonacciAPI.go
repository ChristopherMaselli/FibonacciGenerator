package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type Numbers struct {
	InputNumber int `json:"InputNumber"`
}

type Count struct {
	countNum int
}

func fetchFib(w http.ResponseWriter, r *http.Request) { //Calculate the Fibonacci value of the input number, then record both into Postgres database and return the Fibonacci value.
	var n Numbers

	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var fibNum int = fibNumbers(n.InputNumber)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertDynStmt := "insert into \"Numbers\"(\"InputNumber\",\"FibonacciNumber\") values($1, $2)"
	_, e := db.Exec(insertDynStmt, n.InputNumber, fibNum)
	CheckError(e)

	json.NewEncoder(w).Encode(fibNum)
	return
}

func fetchNum(w http.ResponseWriter, r *http.Request) { //Search for all memorized numbers smaller than the input number and return it

	var n Numbers

	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	selectDynStmt := "Select count(\"FibonacciNumber\") from \"Numbers\" where \"FibonacciNumber\" < $1"
	rows, err := db.Query(selectDynStmt, n.InputNumber)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	snbs := make([]Count, 1)

	for rows.Next() {
		snb := Count{}
		err := rows.Scan(&snb.countNum)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		snbs = append(snbs, snb)
		json.NewEncoder(w).Encode(snb.countNum)
	}
}

func clearData(w http.ResponseWriter, r *http.Request) { //Clear all data in the table
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	_, err = db.Exec("Delete from \"Numbers\"")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func fibNumbers(n int) (x int) { //Recursive function for calculating Fibonacci values of input number
	if n <= 1 {
		x = n
	} else {
		x = fibNumbers(n-1) + fibNumbers(n-2)
	}
	return
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/fib", fetchFib).Methods("POST")
	myRouter.HandleFunc("/fib", fetchNum).Methods("GET")
	myRouter.HandleFunc("/fibClear", clearData)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
