package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciAPIProject(t *testing.T) {
	//Testing fetchFib function
	var testsFetchFib = []struct {
		input    int
		expected int
	}{
		{11, 89},
		{0, 0},
		{1, 1},
		{6, 8},
		{12, 144},
	}

	for _, test := range testsFetchFib {
		results := TestfetchFib(test.input, t)
		assert.Equal(t, test.expected, results)
	}

	//Testing fetchNum function
	var testsFetchNum = []struct {
		input    int
		expected int
	}{
		{120, 9},
		{0, 0},
		{1, 2},
		{50, 7},
		{1000, 12},
	}

	for _, test := range testsFetchNum {
		results := TestfetchNum(test.input, t)
		assert.Equal(t, test.expected, results)
	}
}

func TestfetchFib(input int, t *testing.T) (output int) {
	Str := fmt.Sprintf(`{"InputNumber":%d}`, input)
	var jsonStr = []byte(Str)

	var fibnumT int

	req, err := http.NewRequest("POST", "/fib", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fetchFib)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	json.NewDecoder(rr.Body).Decode(&fibnumT)
	output = fibnumT
	return
}

func TestfetchNum(input int, t *testing.T) (output int) {
	Str := fmt.Sprintf(`{"InputNumber":%d}`, input)
	var jsonStr = []byte(Str)

	var fibnumT int

	req, err := http.NewRequest("GET", "/fib", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fetchNum)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	json.NewDecoder(rr.Body).Decode(&fibnumT)
	output = fibnumT
	return
}
