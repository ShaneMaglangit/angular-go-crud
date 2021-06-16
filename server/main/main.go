package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Status string `json:"status" bson:"status"`
}

type Transaction struct {
	Type   string    `json:"type" bson:"type"`
	Desc   string    `json:"desc" bson:"desc"`
	Amount float32   `json:"amount" bson:"amount"`
	Date   time.Time `json:"date" bson:"date"`
}

var Transactions []Transaction

func main() {
	Transactions = []Transaction{
		{Type: "Income", Desc: "Monthly Earning", Amount: 100.00, Date: time.Now().UTC()},
		{Type: "Expense", Desc: "Bought Pressure Washer", Amount: 50.00, Date: time.Now().UTC()},
	}

	// Setup routing
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/favicon.ico", http.NotFoundHandler())
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/transaction", addTransactionHandler).Methods(http.MethodPost)

	// Get the preferred port to run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s\n", port)
	}

	// Setup CORS
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, cors(router)))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	_, _ = fmt.Fprintf(w, "Server running")
}

func addTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Get request body
	decoder := json.NewDecoder(r.Body)

	// Create new transaction
	var transaction Transaction
	if err := decoder.Decode(&transaction); err != nil {
		http.Error(w, "Failed request", http.StatusBadRequest)
		return
	}
	Transactions = append(Transactions, transaction)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(Response{"ok"})
}
