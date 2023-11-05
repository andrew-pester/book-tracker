package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/andrew-pester/book-tracker/models/books"
	"github.com/gin-gonic/gin"
)

type BookInput struct {
	ISBN        int64     `json:"isbn" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Publisher   string    `json:"publisher" binding:"required"`
	ReleaseTime time.Time `json:"releaseTime" binding:"required"`
}

func AddBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := books.Book{}
	b.ISBN = input.ISBN
	b.Title = input.Title
	b.Author = input.Author
	b.Publisher = input.Publisher
	b.ReleaseTime = input.ReleaseTime
	// b.SaveBook()
	log.Println("Saved Book")
}

func UpdateBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := books.Book{}
	b.ISBN = input.ISBN
	b.Title = input.Title
	b.Author = input.Author
	b.Publisher = input.Publisher
	b.ReleaseTime = input.ReleaseTime
	// b.UpdateBook()
	log.Println("Successfully Saved Book")
}

type InputISBN struct {
	ISBN int64 `json:"isbn" binding:"required"`
}

func GetBookISBN(c *gin.Context) {
	var input InputISBN

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := books.Book{}
	b.ISBN = input.ISBN
	// b.GetBookByISBN()
	log.Println("Successfully retrieved book")
	c.JSON(http.StatusAccepted, gin.H{"Title": b.Title, "Author": b.Author, "releaseTime": b.ReleaseTime.Format(time.RFC3339)})
}

func DeleteBook(c *gin.Context) {
	var input InputISBN

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := books.Book{}
	b.ISBN = input.ISBN
	// b.DeleteBookByISBN()
	log.Println("Successfully deleted book")
}
