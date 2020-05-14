package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, nil)
}

func palindromeNumberHandler(w http.ResponseWriter, r *http.Request) {
}

func organizeBookHandler(w http.ResponseWriter, r *http.Request) {
}

func missingNumberHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/palindrome-number", palindromeNumberHandler)
	http.HandleFunc("/organize-book", organizeBookHandler)
	http.HandleFunc("/missing-number", missingNumberHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
