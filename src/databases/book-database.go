package databases

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/andrew-pester/book-tracker/models/books"
	"github.com/gocql/gocql"
)

var DB *gocql.Session
var DatabaseName string

func init() {
	DatabaseName = os.Getenv("DB_KEYSPACE") + "." + os.Getenv("DB_NAME")
}

func SetupDataBase() {

	host := os.Getenv("DB_HOSTS")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	cluster := gocql.NewCluster(host + ":" + port)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	sess, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	DB = sess
	log.Println("Successfully connected to Database")
	createKeySpace()
	createTable()
	defer DB.Close()

}

func createKeySpace() error {
	query := "CREATE KEYSPACE IF NOT EXISTS " + os.Getenv("DB_KEYSPACE") + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };"
	if err := DB.Query(query).WithContext(context.Background()).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func createTable() error {
	query := "CREATE TABLE IF NOT EXISTS " + os.Getenv("DB_KEYSPACE") + "." + os.Getenv("DB_NAME") + " ( ISBN bigint, title text, author text, publisher text, releaseTime timestamp, PRIMARY KEY (ISBN));"
	if err := DB.Query(query).WithContext(context.Background()).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CreateBook(book *books.Book) (*books.Book, error) {
	ctx := context.Background()

	query := "INSERT INTO " + DatabaseName + " (ISBN, title, author, publisher,releaseTime) VALUES (?, ?, ?, ?, ?)"
	if err := DB.Query(query,
		book.ISBN, book.Title, book.Author, book.Publisher, book.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return book, nil
}
func ReadBookByISBN(book *books.Book) (*books.Book, error) {
	ctx := context.Background()
	query := "SELECT title, author, publisher, releaseTime FROM " + DatabaseName + " WHERE ISBN=" + strconv.FormatInt(book.ISBN, 10)
	if err := DB.Query(query).WithContext(ctx).Scan(&book.Title, &book.Author, &book.Publisher, &book.ReleaseTime); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return book, nil
}

func UpdateBook(book *books.Book) (*books.Book, error) {
	ctx := context.Background()
	query := "UPDATE " + DatabaseName + " SET title=?,author=?,publisher=?,releaseTime=? WHERE ISBN=" + strconv.FormatInt(book.ISBN, 10)
	if err := DB.Query(query, book.Title, book.Author, book.Publisher, book.ReleaseTime).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return book, nil
}

func DeleteBookByISBN(book *books.Book) (*books.Book, error) {
	ctx := context.Background()
	query := "DELETE FROM " + DatabaseName + " WHERE ISBN=" + strconv.FormatInt(book.ISBN, 10)
	if err := DB.Query(query).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return book, nil
}
