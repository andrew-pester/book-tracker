package models

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

var DB *gocql.Session

func ConnectDataBase() {

	host := os.Getenv("DB_HOST")
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

	if err = DB.Query("CREATE KEYSPACE IF NOT EXISTS " + os.Getenv("DB_NAME") + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };").Exec(); err != nil {
		log.Fatal(err)
	}
}
