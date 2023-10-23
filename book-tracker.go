package main

import (
	"github.com/andrew-pester/book-tracker/controllers"
	"github.com/andrew-pester/book-tracker/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//Used to initialize the session for database connection
	models.ConnectDataBase()
	public := r.Group("/api")

	public.POST("/book/add", controllers.AddBook)
	public.DELETE("/book/delete", controllers.DeleteBook)
	public.PATCH("/book/update", controllers.UpdateBook)
	public.GET("/book/get", controllers.GetBookISBN)

	r.Run(":8080")

}
