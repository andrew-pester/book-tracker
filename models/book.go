package models

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
)

type Book struct {
	ISBN        int64
	Title       string
	Author      string
	Publisher   string
	ReleaseTime time.Time
}

var databaseName string = os.Getenv("DB_NAME")

func (b *Book) SaveBook() (*Book, error) {

	ctx := context.Background()
	query := "INSERT INTO " + databaseName + " (ISBN, title, author, publisher,releaseTime) VALUES (?, ?, ?, ?, ?)"
	if err := DB.Query(query,
		b.ISBN, b.Title, b.Author, b.Publisher, b.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) UpdateBook() (*Book, error) {
	ctx := context.Background()
	query := "UPDATE " + databaseName + " SET title=?,author=?,publisher=?,releaseTime=? WHERE ISBN=" + strconv.FormatInt(b.ISBN, 10)
	if err := DB.Query(query, b.Title, b.Author, b.Publisher, b.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) GetBookByISBN() (*Book, error) {
	ctx := context.Background()
	query := "SELECT title, author, publisher, releaseTime FROM " + databaseName + " WHERE ISBN=" + strconv.FormatInt(b.ISBN, 10)
	if err := DB.Query(query).WithContext(ctx).Scan(&b.Title, &b.Author, &b.Publisher, &b.ReleaseTime); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) DeleteBookByISBN() (*Book, error) {
	ctx := context.Background()
	query := "DELETE FROM " + databaseName + " WHERE ISBN=" + strconv.FormatInt(b.ISBN, 10)
	if err := DB.Query(query).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}
