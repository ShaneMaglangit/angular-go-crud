package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type POSTResponse struct {
	Status string    `json:"status" bson:"status"`
	ID     uuid.UUID `json:"id" bson:"id"`
}

type Response struct {
	Status string `json:"status" bson:"status"`
}

type Transaction struct {
	ID     uuid.UUID `json:"id" bson:"id"`
	Type   string    `json:"type" bson:"type"`
	Desc   string    `json:"desc" bson:"desc"`
	Amount float32   `json:"amount" bson:"amount"`
	Date   time.Time `json:"date" bson:"date"`
}

var Transactions []Transaction

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	_, _ = fmt.Fprintf(w, "Server running")
}

func getTransactionHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(Transactions)
}

func addTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Get request body
	decoder := json.NewDecoder(r.Body)

	// Create new transaction.ts
	var transaction Transaction
	if err := decoder.Decode(&transaction); err != nil {
		http.Error(w, "Failed request", http.StatusBadRequest)
		return
	}
	transaction.ID = uuid.New()
	Transactions = append(Transactions, transaction)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(POSTResponse{"created", transaction.ID})
}

func updateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Get request body
	decoder := json.NewDecoder(r.Body)

	// Update transaction
	var transaction Transaction
	if err := decoder.Decode(&transaction); err != nil {
		http.Error(w, "Failed request", http.StatusBadRequest)
		return
	}

	for i, t := range Transactions {
		if t.ID == transaction.ID {
			Transactions[i] = transaction
		}
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response{"updated"})
}

func deleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Get transaction id from params
	transactionId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Failed request", http.StatusBadRequest)
		return
	}

	for i, t := range Transactions {
		if t.ID == transactionId {
			// Remove item from list
			Transactions = append(Transactions[:i], Transactions[i+1:]...)

			// Send response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(Response{"deleted"})
			return
		}
	}

	// Send error response if item is not found
	http.Error(w, "Failed request", http.StatusBadRequest)
}

func main() {
	Transactions = []Transaction{
		{ID: uuid.New(), Type: "Income", Desc: "Monthly Earning", Amount: 100.00, Date: time.Now().UTC()},
		{ID: uuid.New(), Type: "Expense", Desc: "Bought Pressure Washer", Amount: 50.00, Date: time.Now().UTC()},
	}

	// Setup routing
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/favicon.ico", http.NotFoundHandler())
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/transaction", addTransactionHandler).Methods(http.MethodPost)
	router.HandleFunc("/transaction", getTransactionHandler).Methods(http.MethodGet)
	router.HandleFunc("/transaction", updateTransactionHandler).Methods(http.MethodPut)
	router.HandleFunc("/transaction/{id}", deleteTransactionHandler).Methods(http.MethodDelete)

	// Get the preferred port to run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s\n", port)
	}

	// Setup CORS
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, cors(router)))
}
