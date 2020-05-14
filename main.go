package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/reyusfa/golang-test/missingnumber"
	"github.com/reyusfa/golang-test/organizebook"
	"github.com/reyusfa/golang-test/palindromenumber"
)

type missingNumberResponse struct {
	Success int
	Input string
	Result []int
}

type organizeBookResponse struct {
	Success int
	Input string
	Result string
}

type palindromeNumberResponse struct {
	Success int
	Input string
	Result int
}

type errResponse struct {
	Success int
	Message string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if t, err := template.ParseFiles("index.html"); err == nil {
		t.Execute(w, nil)
	}
}

func missingNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()

	if _, ok := query["input"]; ok && r.Method == "GET" {
		input := query["input"][0]

		if data, err := missingnumber.FindMissingNumber(input); err == nil {
			dataJson, _ := json.Marshal(&missingNumberResponse{
				Success: 1,
				Input: input,
				Result: data,
			})

			w.Write(dataJson)
		} else {
			errJson, _ := json.Marshal(&errResponse{
				Success: 0,
				Message: err.Error(),
			})

			w.Write(errJson)
		}
	}
}

func organizeBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()

	if _, ok := query["input"]; ok && r.Method == "GET" {
		input := query["input"][0]

		if data, err := organizebook.OrganizeBooks(input); err == nil {
			dataJson, _ := json.Marshal(&organizeBookResponse{
				Success: 1,
				Input: input,
				Result: data,
			})

			w.Write(dataJson)
		} else {
			errJson, _ := json.Marshal(&errResponse{
				Success: 0,
				Message: err.Error(),
			})

			w.Write(errJson)
		}
	}
}

func palindromeNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()

	if _, ok := query["input"]; ok && r.Method == "GET" {
		input := query["input"][0]

		if data, err := palindromenumber.CountPalindromes(input); err == nil {
			dataJson, _ := json.Marshal(&palindromeNumberResponse{
				Success: 1,
				Input: input,
				Result: data,
			})

			w.Write(dataJson)
		} else {
			errJson, _ := json.Marshal(&errResponse{
				Success: 0,
				Message: err.Error(),
			})

			w.Write(errJson)
		}
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/palindrome-number", palindromeNumberHandler)
	http.HandleFunc("/organize-book", organizeBookHandler)
	http.HandleFunc("/missing-number", missingNumberHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
