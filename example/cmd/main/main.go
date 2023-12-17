package main

import (
	"hello/pkg/routes"

	"gofr.dev/gofr"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := gofr.New()
	db, err := gorm.Open("mysql", "your_connection_string")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	app.Container.MustSet(db)
	routes.RegisterBookStoreRoutes(app)
	app.Start()
}
