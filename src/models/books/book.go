package books

import (
	"time"
)

type Book struct {
	ISBN        int64
	Title       string
	Author      string
	Publisher   string
	ReleaseTime time.Time
}

func New(isbn int64, title string, author string, publisher string, releaseTime time.Time) *Book {
	return &Book{isbn, title, author, publisher, releaseTime}
}

func (b *Book) GetISBN() int64 {
	return b.ISBN
}
func (b *Book) GetTitle() string {
	return b.Title
}
func (b *Book) GetAuthor() string {
	return b.Author
}
func (b *Book) GetPublisher() string {
	return b.Publisher
}
func (b *Book) GetReleaseTime() time.Time {
	return b.ReleaseTime
}
