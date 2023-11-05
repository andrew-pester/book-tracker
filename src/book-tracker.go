package main

import (
	"github.com/andrew-pester/book-tracker/controllers"
	"github.com/andrew-pester/book-tracker/databases"
	"github.com/andrew-pester/book-tracker/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//Used to initialize the session for database connection
	databases.SetupDataBase()
	public := r.Group("/api")
	public.GET("/book/get", controllers.GetBookISBN)
	protected := r.Group("/api/admin")
	protected.Use(middleware.AuthMiddleware())
	protected.DELETE("/book/delete", controllers.DeleteBook)
	protected.PATCH("/book/update", controllers.UpdateBook)
	protected.POST("/book/add", controllers.AddBook)

	r.Run(":8080")

}
