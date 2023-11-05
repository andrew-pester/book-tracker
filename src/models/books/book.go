package books

import (
	"errors"
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

func New(isbn int64, title string, author string, publisher string, releaseTime time.Time) (*Book, error) {

	return &Book{}, nil
}
func validateISBN(isbn int64) error {
	if len(strconv.FormatInt(isbn, 10)) != 13 {
		return errors.New("ISBN is not the approriate length")
	}
	return nil
}
func validateTitle(title string) error {
	return nil
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
func (b *Book) SetISBN(isbn int64) {
	b.ISBN = isbn
}
func (b *Book) SetTitle(title string) {
	b.Title = title
}
func (b *Book) SetAuthor(author string) {
	b.Author = author
}
func (b *Book) SetPublisher(publisher string) {
	b.Publisher = publisher
}
func (b *Book) SetReleaseTime(releaseTime time.Time) {
	b.ReleaseTime = releaseTime
}
