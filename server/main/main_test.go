package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddTransactionHandler(t *testing.T) {
	type test struct {
		name  string
		input []byte
		want  int
	}

	tests := []test{
		{"Complete Field", []byte(`{"type": "Income", "desc": "Weekly earning", "amount": 20.00}`), http.StatusCreated},
		{"Missing Type", []byte(`{"desc": "Weekly earning", "amount": 20.00}`), http.StatusBadRequest},
		{"Missing Amount", []byte(`{"type": "Income", "desc": "Weekly earning"}`), http.StatusBadRequest},
		{"Empty Body", []byte(``), http.StatusBadRequest},
	}

	for _, testCase := range tests {
		req, err := http.NewRequest("POST", "localhost:8080/transaction", bytes.NewReader(testCase.input))
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(addTransactionHandler)
		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != testCase.want {
			t.Fatalf("%s: Handler returned wrong status code, got %v want %v", testCase.name, status, testCase.want)
		}
	}
}
