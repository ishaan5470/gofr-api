package routes

import (
	"gofr.dev/gofr"

	"hello/pkg/controllers"
)

func RegisterBookStoreRoutes(app *gofr.App) {
	app.POST("/book/", controllers.CreateBook)
	app.GET("/book/", controllers.GetBook)
	app.GET("/book/:bookId", controllers.GetBookById)
	app.PUT("/book/:bookId", controllers.UpdateBook)
	app.DELETE("/book/:bookId", controllers.DeleteBook)
}
