package config

import (
	"gofr.dev/gofr"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const connString = "ishaan:ishaan5470@/tablee?charset=utf8&parseTime=True&loc=Local"

func SetupDatabase(app *gofr.App) error {

	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return err
	}

	app.Container.MustSet(db)

	return nil
}
