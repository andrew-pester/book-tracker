package models

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

type Book struct {
	Title       string
	Author      string
	ReleaseTime time.Time
}

func (b *Book) SaveBook() (*Book, error) {

	ctx := context.Background()

	if err := DB.Query("CREATE TABLE IF NOT EXISTS inventory.book ( id UUID, title text, author text, releaseTime timestamp, PRIMARY KEY (title));").WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}

	log.Printf("DB: %s", strconv.FormatBool(DB.Closed()))
	if err := DB.Query("INSERT INTO inventory.book (id, title, author, releaseTime) VALUES (?, ?, ?, ?)",
		gocql.TimeUUID(), b.Title, b.Author, b.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) UpdateBook() (*Book, error) {
	ctx := context.Background()

	if err := DB.Query("UPDATE inventory.book  SET author=?,releaseTime=? WHERE title='"+b.Title+"'",
		b.Author, b.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) GetBook() (*Book, error) {
	ctx := context.Background()
	if err := DB.Query("SELECT author, title, releaseTime from inventory.book WHERE title='"+b.Title+"'").WithContext(ctx).Scan(&b.Author, &b.Title, &b.ReleaseTime); err != nil {
		log.Fatal(err)
	}
	return b, nil
}

func (b *Book) DeleteBook() (*Book, error) {
	ctx := context.Background()
	if err := DB.Query("DELETE FROM inventory.book WHERE title='" + b.Title + "'").WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}
	return b, nil
}
