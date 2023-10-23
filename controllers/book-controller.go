package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/andrew-pester/book-tracker/models"
	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ReleaseTime string `json:"releaseTime" binding:"required"`
}

func AddBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := models.Book{}
	b.Author = input.Author
	b.Title = input.Title
	if inputTime, err := time.Parse(time.RFC3339, input.ReleaseTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		b.ReleaseTime = inputTime
	}
	b.SaveBook()
	log.Println("Saved Book")
}

func UpdateBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := models.Book{}
	b.Author = input.Author
	b.Title = input.Title
	if inputTime, err := time.Parse(time.RFC3339, input.ReleaseTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		b.ReleaseTime = inputTime
	}
	b.UpdateBook()
	log.Println("Saved Book")
}

func GetBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := models.Book{}
	b.Author = input.Author
	b.Title = input.Title
	if inputTime, err := time.Parse(time.RFC3339, input.ReleaseTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		b.ReleaseTime = inputTime
	}
	b.GetBook()
	log.Println("Successfully retrieved book")
	c.JSON(http.StatusAccepted, gin.H{"Title": b.Title, "Author": b.Author, "releaseTime": b.ReleaseTime.Format(time.RFC3339)})
}

func DeleteBook(c *gin.Context) {
	var input BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := models.Book{}
	b.Author = input.Author
	b.Title = input.Title
	if inputTime, err := time.Parse(time.RFC3339, input.ReleaseTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		b.ReleaseTime = inputTime
	}
	b.DeleteBook()
	log.Println("Successfully deleted book")
}
