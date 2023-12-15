package main

import (
	"fmt"
	"log"
	"net/http"
)

//function that hits when we hit the homepage url

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Testing the app")
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hitting the homepage endpoint")

}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
