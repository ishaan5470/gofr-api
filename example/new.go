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

//now setting up the database the schema and create a diffe method and has a database called user and create a user page method and called a mysql databse powered by docker

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
