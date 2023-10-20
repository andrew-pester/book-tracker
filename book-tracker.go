package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/add/:book", addBook)
	public.DELETE("/remove/:book", removeBook)
	public.PATCH("/update/:book", updateBook)
	public.GET("/get/:book", getBook)

	r.Run(":8080")

}

func addBook(c *gin.Context) {

}

func removeBook(c *gin.Context) {

}
func updateBook(c *gin.Context) {

}
func getBook(c *gin.Context) {

}
