package main

import (
	"Practice7/handlers"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	r := gin.Default()

	r.GET("/books", handlers.GetAllBooks)
	r.POST("/books", handlers.AddBooks)
	r.PUT("/books:id", handlers.UpdateBooks)
	r.DELETE("/books:id", handlers.DeleteBook)

	r.Run("localhost:8083")

}
