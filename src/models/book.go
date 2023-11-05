package models

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/andrew-pester/book-tracker/databases"
)

type Book struct {
	ISBN        int64
	Title       string
	Author      string
	Publisher   string
	ReleaseTime time.Time
}

var DatabaseName string

func init() {
	DatabaseName = os.Getenv("DB_KEYSPACE") + "." + os.Getenv("DB_NAME")
}

func (b *Book) SaveBook() (*Book, error) {

	ctx := context.Background()
	query := "INSERT INTO " + DatabaseName + " (ISBN, title, author, publisher,releaseTime) VALUES (?, ?, ?, ?, ?)"
	if err := databases.DB.Query(query,
		b.ISBN, b.Title, b.Author, b.Publisher, b.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) UpdateBook() (*Book, error) {
	ctx := context.Background()
	query := "UPDATE " + DatabaseName + " SET title=?,author=?,publisher=?,releaseTime=? WHERE ISBN=" + strconv.FormatInt(b.ISBN, 10)
	if err := databases.DB.Query(query, b.Title, b.Author, b.Publisher, b.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) GetBookByISBN() (*Book, error) {
	ctx := context.Background()
	query := "SELECT title, author, publisher, releaseTime FROM " + DatabaseName + " WHERE ISBN=" + strconv.FormatInt(b.ISBN, 10)
	if err := databases.DB.Query(query).WithContext(ctx).Scan(&b.Title, &b.Author, &b.Publisher, &b.ReleaseTime); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) DeleteBookByISBN() (*Book, error) {
	ctx := context.Background()
	query := "DELETE FROM " + DatabaseName + " WHERE ISBN=" + strconv.FormatInt(b.ISBN, 10)
	if err := databases.DB.Query(query).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}
