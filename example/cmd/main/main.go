package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// creating the struct first according to the migration file. Till now the migration file consist of table has id and name and so will how struct will reflect
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//function that hits when we hit the homepage url

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Testing the app")
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hitting the homepage endpoint")
}

func getUsers() []*User {
	//first initalize the database connection

	db, err := sql.Open("mysql", "test_user:hello@tcp(db:3306)/test_database")

	//check it error with db connection
	if err != nil {
		panic(err)

	}
	defer db.Close()

	//to grab out results  execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	//otherwise we are gonna loop through the results assigned to user variable
	var users []*User
	for results.Next() {
		var u User
		//scan each result
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, &u)
	}
	return users

}

// users page
func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	fmt.Println("Hitting the user page endpoint")
	json.NewEncoder(w).Encode(users)
}

//now setting up the database the schema and create a diffe method and has a database called user and create a user page method and called a mysql databse powered by docker

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
