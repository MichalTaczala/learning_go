package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform () %v", err)
		return
	}
	fmt.Fprintf(w, "udalo sie")
	name := r.FormValue("name")
	fmt.Fprintf(w, "udalo sie %v", name)

	// if r.URL.Path != "/hello" {
	// 	http.Error(w, "404 not found", http.StatusNotFound)
	// 	return
	// }
	// if r.Method != "GET" {
	// 	http.Error(w, "method not suppoerted", http.StatusNotFound)
	// 	return
	// }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not suppoerted", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello12")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
