package main

import (
	"log"

	"github.com/andrew-pester/book-tracker/controllers"
	"github.com/andrew-pester/book-tracker/middleware"
	"github.com/andrew-pester/book-tracker/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	//Load all env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	//Used to initialize the session for database connection
	models.ConnectDataBase()
	public := r.Group("/api")
	public.GET("/book/get", controllers.GetBookISBN)
	protected := r.Group("/api/admin")
	protected.Use(middleware.AuthMiddleware())
	protected.DELETE("/book/delete", controllers.DeleteBook)
	protected.PATCH("/book/update", controllers.UpdateBook)
	protected.POST("/book/add", controllers.AddBook)

	r.Run(":8080")

}
