package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/andrew-pester/book-tracker/databases"
	"github.com/andrew-pester/book-tracker/models/books"
	"github.com/gin-gonic/gin"
)

type InputISBN struct {
	ISBN int64 `json:"isbn" binding:"required,isbn13"`
}
type BookInput struct {
	ISBN        int64     `json:"isbn" binding:"required,isbn13"`
	Title       string    `json:"title" binding:"required,alphanum"`
	Author      string    `json:"author" binding:"required,alphanum"`
	Publisher   string    `json:"publisher" binding:"required,alphanum"`
	ReleaseTime time.Time `json:"releaseTime" binding:"required,datetime"`
}

func AddBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := books.New(input.ISBN, input.Title, input.Author, input.Publisher, input.ReleaseTime)
	databases.CreateBook(b)
	log.Println("Saved Book")
}

func UpdateBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := books.New(input.ISBN, input.Title, input.Author, input.Publisher, input.ReleaseTime)
	databases.UpdateBook(b)
	log.Println("Successfully Saved Book")
}

func GetBookISBN(c *gin.Context) {
	var input InputISBN

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := &books.Book{}
	b.ISBN = input.ISBN
	databases.ReadBookByISBN(b)
	log.Println("Successfully retrieved book")
	c.JSON(http.StatusAccepted, gin.H{"Title": b.Title, "Author": b.Author, "releaseTime": b.ReleaseTime.Format(time.RFC3339)})
}

func DeleteBook(c *gin.Context) {
	var input InputISBN

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := &books.Book{}
	b.ISBN = input.ISBN
	databases.DeleteBookByISBN(b)
	log.Println("Successfully deleted book")
}
