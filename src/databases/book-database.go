package databases

import (
	"context"
	"log"
	"os"

	"github.com/gocql/gocql"
)

var DB *gocql.Session

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
